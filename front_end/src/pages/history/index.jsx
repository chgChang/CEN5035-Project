import { CloseCircleOutlined } from '@ant-design/icons';
import { Avatar, Card, Col, Popover, Row, message, List, Typography, Button } from 'antd';
import { useState } from 'react';
import { useRequest } from 'umi';
import ProForm, {
  ProFormDateRangePicker,
  ProFormSelect,
  ProFormText,
  ProFormTimePicker,
} from '@ant-design/pro-form';
import { EditableProTable } from '@ant-design/pro-table';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import styles from './style.less';
const { Paragraph } = Typography;

const AdvancedForm = () => {
  // const onFinish = async (values) => {
  //   setError([]);

  //   try {
  //     await fakeSubmitForm(values);
  //     message.success('提交成功');
  //   } catch {
  //     // console.log
  //   }
  // };

  // const onFinishFailed = (errorInfo) => {
  //   setError(errorInfo.errorFields);
  // };
  const lists = require("./res_getOderHis_sus.json");
  const list = lists?.histories || [];

  const orderList = list && (
    <List
      itemLayout="horizontal"
      dataSource={list}
      renderItem={(item) => (
        <List.Item>
          <Card
            title = {item.orderDate}
            bordered={false}
            className={styles.card}
            style={{ width: '100%' }}
            hoverable
          >
            <Card.Meta
              description={
                <List
                  dataSource={item.itemList}
                  renderItem={(orderItem) => (
                    <List.Item
                      actions={[
                        <Button shape="round" className={styles.addcartbtn}>
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
