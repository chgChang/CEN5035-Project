import { CloseCircleOutlined } from '@ant-design/icons';
import { Card, Col, Popover, Row, message, List } from 'antd';
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
import { fakeSubmitForm } from './service';
import styles from './style.less';

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

  const { data, loading, run } = useRequest((values) => {
    return fakeSubmitForm({
      count: 8,
    });
  });
  console.log(data);
  const list = data?.list || [];

  const cardList = list && (
    <List
      loading={loading}
      itemLayout="horizontal"
      dataSource={list}
      renderItem={(item) => (
        <List.Item>
          <Card
            className={styles.card}
            style={{ height: 350 }}
            hoverable
            cover={
              <img
                style={{ margin: '0 auto', maxHeight: 200, width: 'auto', maxWidth: '100%' }}
                alt={item.itemName}
                src={item.picurl}
              />
            }
          >
            <Card.Meta
              title={<a>{item.itemName}</a>}
              description={
                <Paragraph
                  className={styles.item}
                  ellipsis={{
                    rows: 2,
                  }}
                >
                  {item.description}
                </Paragraph>
              }
            />

            <div className={styles.cardItemContent}>
              <Button shape="round" className={styles.addcartbtn}>
                Add to Cart
              </Button>
            </div>
          </Card>
        </List.Item>
      )}
    />
  );

  return (
    <div>
      <PageContainer content="Welcome back to check your order history~">
        <Card
          className={styles.card}
          bordered={false}
          cover={
            <div>
              <Row>
                <Col span={4}>Order Placed</Col>
                <Col span={4}>Total</Col>
                <Col span={4} offset={15}>
                  OrderID
                </Col>
              </Row>
              <Row>
                <Col span={4}>xx-xx-xxxx</Col>
                <Col span={4}>00.00 </Col>
                <Col span={4} offset={15}>
                  xxxxxxxxxx
                </Col>
              </Row>
            </div>
          }
        >
          {cardList}
        </Card>
      </PageContainer>
    </div>
  );
};

export default AdvancedForm;
