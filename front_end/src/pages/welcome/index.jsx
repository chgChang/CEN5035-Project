import { PageContainer } from '@ant-design/pro-layout';
import { Input, Card, Col, Form, List, Row, Select, Typography, Button, message } from 'antd';
import { useEffect, useState, useRef } from 'react';
import useRequest from '@ahooksjs/use-request';
import { queryItemList, searchItem, add2Cart } from './service';
import styles from './style.less';

const tabList = [
  {
    key: 'All',
    tab: 'All',
  },
  {
    key: 'Apple',
    tab: 'Apple',
  },
];

const { Paragraph } = Typography;

const searchval = "";

const Search = (props) => {

  const addCart = async (id) => {
    const res = await add2Cart({itemid: id, quantity: 1});
    if (res.status === "success") {
      message.success(res.msg);
      return;
    } else {
      message.error(res.msg);
    } 
  };

  const { data, loading, mutate } = useRequest(queryItemList);

  const { run: postRun } = useRequest(
    (params) => {
        return searchItem(params);
    },
    {
      manual: true,
      onSuccess: (result) => {
        console.log("this is result");
        console.log(result);
        mutate(result);
      },
    }
  );


  const handleTabChange = (key) => {
    if (key === 'All') {
      postRun("");
    } else {
      postRun(key);
    }
  };

  const handleFormSubmit = (value) => {
    postRun(value);
  };

  const getTabKey = () => {
    const { match, location } = props;
    const url = match.path === '/' ? '' : match.path;
    const tabKey = location.pathname.replace(`${url}/`, '');

    if (tabKey && tabKey !== '/') {
      return tabKey;
    }

    return 'projects';
  };

  const list = data?.list || [];
  console.log(list);
  const cardList = list && (
    <List
      rowKey="id"
      loading={loading}
      grid={{
        gutter: 16,
        xs: 1,
        sm: 2,
        md: 3,
        lg: 3,
        xl: 4,
        xxl: 4,
      }}
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
                alt={item.name}
                src={item.pic_url}
              />
            }
          >
            <Card.Meta
              title={<a>{item.name}</a>}
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
              <Button shape="round" className={styles.addcartbtn} onClick = {() => addCart(item.id)}>
                Add to Cart
              </Button>
            </div>
          </Card>
        </List.Item>
      )}
    />
  );
  const formItemLayout = {
    wrapperCol: {
      xs: {
        span: 24,
      },
      sm: {
        span: 16,
      },
    },
  };

  return (
    <PageContainer
      content={
        <div
          style={{
            textAlign: 'center',
          }}
        >
          <Input.Search
            placeholder="Please input"
            enterButton="Search"
            size="large"
            onSearch={handleFormSubmit}
            style={{
              maxWidth: 522,
              width: '100%',
            }}
          />
        </div>
      }
      tabList={tabList}
      tabActiveKey={getTabKey()}
      onTabChange={handleTabChange}
    >
      <div className={styles.coverCardList}>
        <div className={styles.cardList}>{cardList}</div>
      </div>
    </PageContainer>
  );
};

export default Search;
