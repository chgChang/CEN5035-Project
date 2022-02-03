import { AudioTwoTone } from '@ant-design/icons';
import { Card, Col, Form, List, Row, Select, Typography, Button } from 'antd';
import moment from 'moment';
import { useRequest } from 'umi';
import AvatarList from './components/AvatarList';
import { queryItemList, queryFakeList } from './service';
import styles from './style.less';
const { Option } = Select;
const FormItem = Form.Item;
const { Paragraph } = Typography;

const getKey = (id, index) => `${id}-${index}`;

const Projects = () => {
  const { data, loading, run } = useRequest((values) => {
    return queryItemList({
      count: 8,
    });
  });
  console.log(data);
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
              {/* <span>{moment(item.updatedAt).fromNow()}</span>
              <div className={styles.avatarList}>
                <AvatarList size="small">
                  {item.members.map((member, i) => (
                    <AvatarList.Item
                      key={getKey(item.id, i)}
                      src={member.avatar}
                      tips={member.name}
                    />
                  ))}
                </AvatarList>
              </div> */}
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
