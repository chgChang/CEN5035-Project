import {
  AmazonOutlined,
  LockOutlined,
  MailOutlined,
  MobileOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { Form, Alert, message, Tabs, Popover, Progress, Input } from "antd";
import React, { useState } from "react";
import {
  ProFormCaptcha,
  ProFormCheckbox,
  ProFormText,
  ProForm,
  LoginForm,
  Submit,
} from "@ant-design/pro-form";
import { useIntl, history, FormattedMessage, SelectLang, useModel } from "umi";
import Footer from "@/components/Footer";
import { login, register } from "@/services/ant-design-pro/api";
import { getFakeCaptcha } from "@/services/ant-design-pro/login";
import { fakeRegister } from "./service";
import styles from "./index.less";

const LoginMessage = ({ content }) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

const passwordProgressMap = {
  ok: "success",
  pass: "normal",
  poor: "exception",
};

const Login = () => {
  const [userLoginState, setUserLoginState] = useState({});
  const [type, setType] = useState("account");
  const [visible, setVisible] = useState(false);
  const { initialState, setInitialState } = useModel("@@initialState");
  const intl = useIntl();
  const [popover, setPopover] = useState(false);
  const confirmDirty = false;
  const [form] = Form.useForm();

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();

    if (userInfo) {
      await setInitialState((s) => ({ ...s, currentUser: userInfo }));
    }
  };

  const handleSubmit = async (values) => {
    if (type === "account") {
      try {
        // 登录
        const res = await login({ ...values, type });
  
        if (res.status === "success") {
          const defaultLoginSuccessMessage = intl.formatMessage({
            id: "pages.login.success",
            defaultMessage: "登录成功！",
          });
          message.success(defaultLoginSuccessMessage);
          await fetchUserInfo();
          /** 此方法会跳转到 redirect 参数所在的位置 */
  
          if (!history) return;
          const { query } = history.location;
          const { redirect } = query;
          history.push(redirect || "/");
          return;
        } else {
          message.error(res.msg);
        } 
  
        console.log(res); // 如果失败去设置用户错误信息
  
        setUserLoginState(res);
      } catch (error) {
        const defaultLoginFailureMessage = intl.formatMessage({
          id: "pages.login.failure",
          defaultMessage: "登录失败，请重试！",
        });
        message.error(defaultLoginFailureMessage);
      }
    } else {
      try {
        // 注册
        const reginfo = {
          email: values.email,
          username: values.username,
          password: values.registerPassword,
        };
        const res = await register({ ...reginfo, type });
  
        if (res.status === "success") {
          const defaultLoginSuccessMessage = intl.formatMessage({
            id: "pages.register.success",
            defaultMessage: "注册成功！",
          });
          message.success(defaultLoginSuccessMessage);
          const loginfo = {
            email: values.email,
            password: values.registerPassword,
          };
          const res = await login({ ...loginfo, type });
          if (res.status === "success") {
            const defaultLoginSuccessMessage = intl.formatMessage({
              id: "pages.login.success",
              defaultMessage: "登录成功！",
            });
            message.success(defaultLoginSuccessMessage);
            await fetchUserInfo();
            /** 此方法会跳转到 redirect 参数所在的位置 */
    
            if (!history) return;
            const { query } = history.location;
            const { redirect } = query;
            history.push(redirect || "/");
          } 
          return;
        } else {
          message.error(res.msg);
        }
  
        console.log(res); // 如果失败去设置用户错误信息
  
        setUserLoginState(res);
      } catch (error) {
        const defaultRegisterFailureMessage = intl.formatMessage({
          id: "pages.register.failure",
          defaultMessage: "注册失败，请重试！",
        });
        message.error(defaultRegisterFailureMessage);
      }
    }
  };

  const { status, type: loginType } = userLoginState;

  const checkPassword = (_, value) => {
    const promise = Promise; // 没有值的情况

    if (!value) {
      setVisible(!!value);
      return promise.reject("Please input Password!");
    } // 有值的情况

    if (!visible) {
      setVisible(!!value);
    }

    setPopover(!popover);

    if (value.length < 6) {
      return promise.reject("");
    }

    if (value && confirmDirty) {
      form.validateFields(["confirm"]);
    }

    return promise.resolve();
  };

  const getPasswordStatus = () => {
    const value = form.getFieldValue("registerPassword");

    if (value && value.length > 9) {
      return "ok";
    }

    if (value && value.length > 5) {
      return "pass";
    }

    return "poor";
  };

  const renderPasswordProgress = () => {
    const value = form.getFieldValue("registerPassword");
    const passwordStatus = getPasswordStatus();
    return value && value.length ? (
      <div className={styles[`progress-${passwordStatus}`]}>
        <Progress
          status={passwordProgressMap[passwordStatus]}
          className={styles.progress}
          strokeWidth={6}
          percent={value.length * 10 > 100 ? 100 : value.length * 10}
          showInfo={false}
        />
      </div>
    ) : null;
  };

  const passwordStatusMap = {
    ok: (
      <div className={styles.success}>
        <span>Safety：Strong</span>
      </div>
    ),
    pass: (
      <div className={styles.warning}>
        <span>Safety：Medium</span>
      </div>
    ),
    poor: (
      <div className={styles.error}>
        <span>Saftey：Short</span>
      </div>
    ),
  };

  const checkConfirm = (_, value) => {
    const promise = Promise;
    
    if (value && value !== form.getFieldValue("registerPassword")) {
      return promise.reject("Two passwords are not same!");
    }

    return promise.resolve();
  };

  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang />}
      </div>
      <div className={styles.content}>
        <LoginForm
          form={form}
          logo={<img style={{height: 38, width: 139}} alt=""  src="/logo.svg" />}
          title={<span style={{marginLeft: 88}}>Gator Amazon</span>}
          subTitle={intl.formatMessage({
            id: "pages.layouts.userLayout.title",
          })}
          initialValues={{
            autoLogin: true,
          }}
          onFinish={async (values) => {
            await handleSubmit(values);
          }}
        >
          {status === "error" && loginType === "account" && (
            <LoginMessage
              content={intl.formatMessage({
                id: "pages.login.accountLogin.errorMessage",
                defaultMessage: "账户或密码错误(admin/ant.design)",
              })}
            />
          )}
          <Tabs activeKey={type} onChange={setType}>
            <Tabs.TabPane
              key="account"
              tab={intl.formatMessage({
                id: "pages.login.accountLogin.tab",
                defaultMessage: "账户密码登录",
              })}
            />
            <Tabs.TabPane
              key="register"
              tab={intl.formatMessage({
                id: "pages.login.phoneLogin.tab",
                defaultMessage: "Register",
              })}
            />
          </Tabs>
          {type === "account" && (
            <>
              <ProFormText
                name="email"
                fieldProps={{
                  size: "large",
                  prefix: <MailOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: "pages.login.emailAddress.placeholder",
                  defaultMessage: "email address",
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.login.emailAddress.required"
                        defaultMessage="请输入email!"
                      />
                    ),
                  },
                  {
                    type: "email",
                    message: (
                      <FormattedMessage
                        id="pages.login.emailAddress.invalid"
                        defaultMessage="Valid email"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText.Password
                name="password"
                fieldProps={{
                  size: "large",
                  prefix: <LockOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: "pages.login.password.placeholder",
                  defaultMessage: "密码: ant.design",
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.login.password.required"
                        defaultMessage="请输入密码！"
                      />
                    ),
                  },
                ]}
              />
            </>
          )}

          {type === "register" && (
            <>
              <ProFormText //Input email address
                fieldProps={{
                  size: "large",
                  prefix: <MailOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: "pages.login.emailAddress.required",
                  defaultMessage: "Input email",
                })}
                name="email"
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.login.emailAddress.required"
                        defaultMessage="Input email"
                      />
                    ),
                  },
                  {
                    type: "email",
                    message: (
                      <FormattedMessage
                        id="pages.login.emailAddress.invalid"
                        defaultMessage="Valid email"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText //Input email address
                fieldProps={{
                  size: "large",
                  prefix: <UserOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: "pages.register.username.required",
                  defaultMessage: "Input email",
                })}
                name="username"
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="Input username"
                      />
                    ),
                  },
                ]}
              />
              <Popover
                getPopupContainer={(node) => {
                  if (node && node.parentNode) {
                    return node.parentNode;
                  }

                  return node;
                }}
                content={
                  visible && (
                    <div
                      style={{
                        padding: "4px 0",
                      }}
                    >
                      {passwordStatusMap[getPasswordStatus()]}
                      {renderPasswordProgress()}
                      <div
                        style={{
                          marginTop: 10,
                        }}
                      >
                        <span>
                          Please use at least 6 characters. Do not use easy
                          password.
                        </span>
                      </div>
                    </div>
                  )
                }
                overlayStyle={{
                  width: 240,
                }}
                placement="right"
                visible={visible}
              >
                <ProFormText.Password
                  fieldProps={{
                    size: "large",
                    prefix: <LockOutlined className={styles.prefixIcon} />,
                  }}
                  placeholder={intl.formatMessage({
                    id: "pages.login.captcha.placeholder",
                    defaultMessage: "至少6位密码，区分大小写",
                  })}
                  name="registerPassword"
                  // className={
                  //   form.getFieldValue('registerPassword') &&
                  //   form.getFieldValue('registerPassword').length > 0 &&
                  //   styles.password
                  // }
                  rules={[
                    {
                      validator: checkPassword,
                    },
                  ]}
                />
              </Popover>
              <ProFormText.Password
                fieldProps={{
                  size: "large",
                  prefix: <LockOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: "pages.login.captcha.placeholder",
                  defaultMessage: "Input password!",
                })}
                name="confirmPassword"
                rules={[
                  {
                    required: true,
                    message: "Retype password",
                  },
                  {
                    validator: checkConfirm,
                  },
                ]}
              />
            </>
          )}
          <div
            style={{
              marginBottom: 24,
            }}
          >
            <ProFormCheckbox noStyle name="autoLogin">
              <FormattedMessage
                id="pages.login.rememberMe"
                defaultMessage="自动登录"
              />
            </ProFormCheckbox>
            {/* <Submit>Submit</Submit> */}
          </div>
        </LoginForm>
      </div>
      <Footer />
    </div>
  );
};

export default Login;
