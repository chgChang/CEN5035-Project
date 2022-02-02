import { Card, Col, Form, List, Row, Select, Typography } from 'antd';
import moment from 'moment';
import { useRequest } from 'umi';
import AvatarList from './components/AvatarList';
import { queryFakeList } from './service';
import styles from './style.less';
const { Option } = Select;
const FormItem = Form.Item;
const { Paragraph } = Typography;

const getKey = (id, index) => `${id}-${index}`;

const Projects = () => {
  const { data, loading, run } = useRequest((values) => {
    console.log('form data', values);
    return queryFakeList({
      count: 8,
    });
  });
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
          <Card className={styles.card} hoverable cover={<img alt={item.title} src={item.cover} />}>
            <Card.Meta
              title={<a>{item.title}</a>}
              description={
                <Paragraph
                  className={styles.item}
                  ellipsis={{
                    rows: 2,
                  }}
                >
                  {item.subDescription}
                </Paragraph>
              }
            />
            {/* <div className={styles.cardItemContent}>
              <span>{moment(item.updatedAt).fromNow()}</span>
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
              </div>
            </div> */}
          </Card>
        </List.Item>
      )}
    />
  );

  return (
    <div className={styles.coverCardList}>
      <div className={styles.cardList}>{cardList}</div>
    </div>
  );
};

export default Projects;
