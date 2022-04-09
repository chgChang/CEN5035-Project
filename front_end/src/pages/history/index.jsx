import { Avatar, Card, Col, Popover, Row, message, List, Typography, Button, Drawer } from 'antd';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import ProDescriptions from '@ant-design/pro-descriptions';
import { useState, useRef } from 'react';
import styles from './style.less';
import useRequest from '@ahooksjs/use-request';
import { getOrderHis, add2Cart } from './service';
const { Paragraph } = Typography;

const AdvancedForm = () => {
  const [showDetail, setShowDetail] = useState(false);
  const [curId, setcurId] = useState("");
  const [curName, setcurName] = useState("");
  const [curDes, setcurDes] = useState("");
  const [curPrice, setcurPrice] = useState("");
  const [curPic, setcurPic] = useState("");

  const { data, loading, mutate } = useRequest(getOrderHis);
  const list = data?.histories || [];

  const addCart = async (id) => {
    const res = await add2Cart({itemid: id, quantity: 1});
    if (res.status === "success") {
      message.success(res.msg);
      return;
    } else {
      message.error(res.msg);
    } 
  };

  const extraContent = (item) => (
    <div className={styles.extraContent}>
      {/* <div style={{fontSize: 16, fontWeight: 600, float: 'left'}}> Order#: {item.orderId} </div> */}
      <div style={{fontSize: 16, fontWeight: 600}}> Order Date: {item.orderDate} </div>
    </div>
  );

  const orderList = list && (
    <List
      itemLayout="horizontal"
      loading={loading}
      dataSource={list}
      renderItem={(item) => (
        <List.Item>
          <Card
            title = {"Order#: " + item.orderId}
            bordered={false}
            className={styles.card}
            style={{ width: '100%' }}
            extra={extraContent(item)}
            hoverable
          >
            <Card.Meta
              description={
                <List
                  dataSource={item.itemList}
                  renderItem={(orderItem) => (
                    <List.Item
                      actions={[
                        <Button shape="round" onClick = {() => addCart(orderItem.itemId)}>
                          Add to Cart
                        </Button>,
                      ]}
                    >
                      <List.Item.Meta
                        avatar={
                          <Avatar src={orderItem.picUrl} shape="square" size="large" />
                        }
                        title={<a onClick={() => {
                          setcurId(orderItem.id);
                          setcurName(orderItem.itemName);
                          setcurPrice(orderItem.price);
                          setcurDes(orderItem.description);
                          setcurPic(orderItem.picUrl);
                          setShowDetail(true);
                        }}>{orderItem.itemName}</a>}
                        description={orderItem.description}
                      />
                    </List.Item>
                  )}
                />
              }
            />
            <Card.Meta
              description={<div style={{ alignItems: 'flex-end' }}>Total price: ${item.totalPrice}</div>}
            />
          </Card>
        </List.Item>
      )}
    />
  );

  return (
    <div>
      <PageContainer content="Welcome back to check your order history~">
        {orderList}
        <Drawer
          width={800}
          visible={showDetail}
          onClose={() => {
            setcurId("");
            setcurName("");
            setcurPrice("");
            setcurDes("");
            setcurPic("");
            setShowDetail(false);
         }}
          closable={false}
        >
          {curName && (
            <ProDescriptions
              column={1}
              title={curName}
              // request={async () => ({
              //   data: cItem || {},
              // })}
              params={{
                id: curName,
              }}
            >
              <ProDescriptions.Item>
                <p className={styles.pPic}>
                  <img
                    className={styles.drawPic}
                    alt={curName}
                    src={curPic}
                  />
                </p>
              </ProDescriptions.Item>
              <ProDescriptions.Item dataIndex="price" label="Price" valueType="price">
                {curPrice}
              </ProDescriptions.Item>
              <ProDescriptions.Item dataIndex="description" label="Description" valueType="textarea">
                {curDes}
              </ProDescriptions.Item>
              <ProDescriptions.Item>
                <Button shape="round" className={styles.addcartbtn} onClick = {() => addCart(curId)}>
                  Add to Cart
                </Button>
              </ProDescriptions.Item>
            </ProDescriptions>
          )}
        </Drawer>
      </PageContainer>
    </div>
  );
};

export default AdvancedForm;
