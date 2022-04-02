import { PageContainer } from '@ant-design/pro-layout';
import { Input } from 'antd';
import { history } from 'umi';
const tabList = [
  {
    key: 'All',
    tab: 'All',
  },
  {
    key: 'Apple',
    tab: 'Apple',
  },
  // {
  //   key: 'Clothing',
  //   tab: 'Clothing',
  // },
  // {
  //   key: 'Shoes',
  //   tab: 'Shoes',
  // },
  // {
  //   key: 'Books',
  //   tab: 'Books',
  // },
  // {
  //   key: 'Electronics',
  //   tab: 'Electronics',
  // },
];

const searchval = "";

const Search = (props) => {
  const handleTabChange = (key) => {
    const { match } = props;
    const url = match.url === '/' ? '' : match.url;
    if (key === 'All') {
      history.push(`${url}/Items`);
    } else {
      console.log(key);
      history.push(`${url}/search/${key}`);
      
    }
  };

  const handleFormSubmit = (value) => {
    // search function
    console.log(value);
    const { match } = props;
    const url = match.url === '/' ? '' : match.url;
    // window.location.replace(`${url}/search/${value}`);
    history.push(`${url}/search/${value}`);
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
      {props.children}
    </PageContainer>
  );
};

export default Search;
