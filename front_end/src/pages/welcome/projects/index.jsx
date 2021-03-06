import { AudioTwoTone } from '@ant-design/icons';
import { Card, Col, Form, List, Row, Select, Typography, Button, message } from 'antd';
import moment from 'moment';
import useRequest from '@ahooksjs/use-request';
import AvatarList from './components/AvatarList';
import { queryItemList, add2Cart } from './service';
import styles from './style.less';

const { Paragraph } = Typography;

const Projects = (props) => {

  const addCart = async (id) => {
    const res = await add2Cart({itemid: id, quantity: 1});
    if (res.status === "success") {
      message.success(res.msg);
      return;
    } else {
      message.error(res.msg);
    } 
  };

  const { data, loading, run } = useRequest(queryItemList);
  const list = data?.list || [];
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
    <div className={styles.coverCardList}>
      <div className={styles.cardList}>{cardList}</div>
    </div>
  );
};

export default Projects;
