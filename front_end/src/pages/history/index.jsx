import { Avatar, Card, Col, Popover, Row, message, List, Typography, Button } from 'antd';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import styles from './style.less';
import useRequest from '@ahooksjs/use-request';
import { getOrderHis, add2Cart } from './service';
const { Paragraph } = Typography;

const AdvancedForm = () => {
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
                        <Button shape="round" className={styles.addcartbtn} onClick = {() => addCart(orderItem.itemId)}>
                          Add to Cart
                        </Button>,
                      ]}
                    >
                      <List.Item.Meta
                        avatar={
                          <Avatar src={orderItem.picUrl} shape="square" size="large" />
                        }
                        title={<a>{orderItem.itemName}</a>}
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
      </PageContainer>
    </div>
  );
};

export default AdvancedForm;
