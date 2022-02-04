import { PageContainer } from '@ant-design/pro-layout';
import { Input } from 'antd';
import { history } from 'umi';
const tabList = [
  {
<<<<<<< HEAD
    key: 'All',
    tab: 'All',
  },
  {
    key: 'Clothing/Shoes',
    tab: 'Clothing/Shoes',
  },
  {
    key: 'Books',
    tab: 'Books',
  },
  {
    key: 'Movie/Music/Games',
    tab: 'Movie/Music/Games',
  },
  {
    key: 'Electronics',
    tab: 'Electronics',
=======
    key: 'articles',
    tab: '文章',
  },
  {
    key: 'projects',
    tab: '项目',
  },
  {
    key: 'applications',
    tab: '应用',
>>>>>>> changzhou
  },
];

const Search = (props) => {
  const handleTabChange = (key) => {
    const { match } = props;
    const url = match.url === '/' ? '' : match.url;

    switch (key) {
<<<<<<< HEAD
      case 'Clothing/Shoes':
        history.push(`${url}/applications`);
        break;

      case 'All':
        history.push(`${url}/projects`);
        break;

      case 'Books':
        history.push(`${url}/applications`);
        break;

      case 'Movie/Music/Games':
        history.push(`${url}/applications`);
        break;

      case 'Electronics':
        history.push(`${url}/applications`);
=======
      case 'articles':
        history.push(`${url}/articles`);
        break;

      case 'applications':
        history.push(`${url}/applications`);
        break;

      case 'projects':
        history.push(`${url}/projects`);
>>>>>>> changzhou
        break;

      default:
        break;
    }
  };

  const handleFormSubmit = (value) => {
<<<<<<< HEAD
    // search function

=======
    // eslint-disable-next-line no-console
>>>>>>> changzhou
    console.log(value);
  };

  const getTabKey = () => {
    const { match, location } = props;
    const url = match.path === '/' ? '' : match.path;
    const tabKey = location.pathname.replace(`${url}/`, '');

    if (tabKey && tabKey !== '/') {
      return tabKey;
    }

<<<<<<< HEAD
    return 'projects';
=======
    return 'articles';
>>>>>>> changzhou
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
            placeholder="请输入"
            enterButton="搜索"
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
      {props.children}
    </PageContainer>
  );
};

export default Search;
