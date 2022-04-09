import React from 'react';
import { UploadOutlined } from '@ant-design/icons';
import { Button, Input, Upload, message } from 'antd';
import ProForm, {
  ProFormDependency,
  ProFormFieldSet,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-form';
import useRequest from '@ahooksjs/use-request';
import { getcurrentUser } from '../service';
import styles from './BaseView.less';


const AvatarView = ({ avatar }) => (
  <>
    <div className={styles.avatar_title}>Image</div>
    <div className={styles.avatar}>
      <img src={avatar} alt="avatar" />
    </div>
    {/* <Upload showUploadList={false}>
      <div className={styles.button_view}>
        <Button>
          <UploadOutlined />
          Update Image
        </Button>
      </div>
    </Upload> */}
  </>
);

const BaseView = () => {
  const { data:currentUser, loading } = useRequest(() => {
    return getcurrentUser();
  });


  console.log(currentUser);
  console.log(currentUser?.data || []);

  // const getAvatarURL = () => {
  //   if (currentUser) {
  //     if (currentUser.avatar) {
  //       return currentUser.avatar;
  //     }

  //     const url = 'https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png';
  //     return url;
  //   }

  //   return '';
  // };

  const handleFinish = async () => {
    message.success('Update Successfully!');
  };

  return (
    <div className={styles.baseView}>
      {loading ? null : (
        <>
          <div className={styles.left}>
            <ProForm
              layout="vertical"
              onFinish={handleFinish}
              submitter={{
                resetButtonProps: {
                  style: {
                    display: 'none',
                  },
                },
                submitButtonProps: {
                  children: 'Update',
                },
              }}
              initialValues={{...currentUser.data}}
              hideRequiredMark
            >
              <ProFormText
                width="md"
                name="email"
                label="Email"
                rules={[
                  {
                    required: true,
                    message: 'Please input your email address!',
                  },
                ]}
              />
              <ProFormText
                width="md"
                name="username"
                label="Name"
                rules={[
                  {
                    required: true,
                    message: 'please input your name!',
                  },
                ]}
              />
              
              {/* <ProFormText
                width="md"
                name="password"
                label="Password"
                rules={[
                  {
                    required: true,
                    message: 'please input your password!',
                  },
                ]}
              /> */}
         
            </ProForm>
          </div>
          <div className={styles.right}>
            <AvatarView avatar={'https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png'} />
          </div>
        </>
      )}
    </div>
  );
};

export default BaseView;
