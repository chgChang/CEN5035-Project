import { useIntl } from 'umi';
import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-layout';
export default () => {
  const intl = useIntl();
<<<<<<< HEAD
  const currentYear = new Date().getFullYear();
  return (
    <DefaultFooter
      copyright={`${currentYear} Designed by Wei Wu, Chang Zhou, JiaNan He, Chi Zhang`}
      links={[
        {
          key: 'OurGithub',
          title: 'Our Github Link',
          href: 'https://github.com/chgChang/CEN5035-Project',
=======
  const defaultMessage = intl.formatMessage({
    id: 'app.copyright.produced',
    defaultMessage: '蚂蚁集团体验技术部出品',
  });
  const currentYear = new Date().getFullYear();
  return (
    <DefaultFooter
      copyright={`${currentYear} ${defaultMessage}`}
      links={[
        {
          key: 'Ant Design Pro',
          title: 'Ant Design Pro',
          href: 'https://pro.ant.design',
>>>>>>> changzhou
          blankTarget: true,
        },
        {
          key: 'github',
          title: <GithubOutlined />,
          href: 'https://github.com/ant-design/ant-design-pro',
          blankTarget: true,
        },
        {
<<<<<<< HEAD
          key: 'Ant Design Pro',
          title: 'Ant Design Pro',
          href: 'https://pro.ant.design',
=======
          key: 'Ant Design',
          title: 'Ant Design',
          href: 'https://ant.design',
>>>>>>> changzhou
          blankTarget: true,
        },
      ]}
    />
  );
};
