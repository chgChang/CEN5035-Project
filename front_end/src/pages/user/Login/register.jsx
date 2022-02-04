import { useState, useEffect } from "react";
import {
  Form,
  Button,
  Col,
  Input,
  Popover,
  Progress,
  Row,
  Select,
  message,
} from "antd";
import { Link, useRequest, history } from "umi";
import { fakeRegister } from "./service";
import styles from "./style.less";
const FormItem = Form.Item;
const { Option } = Select;
const InputGroup = Input.Group;
const passwordProgressMap = {
  ok: "success",
  pass: "normal",
  poor: "exception",
};

const Register = () => {
  const [count, setCount] = useState(0);
  const [prefix, setPrefix] = useState("86");
  const confirmDirty = false;
  let interval;
  const [form] = Form.useForm();
  useEffect(
    () => () => {
      clearInterval(interval);
    },
    [interval]
  );

  const onGetCaptcha = () => {
    let counts = 59;
    setCount(counts);
    interval = window.setInterval(() => {
      counts -= 1;
      setCount(counts);

      if (counts === 0) {
        clearInterval(interval);
      }
    }, 1000);
  };

  const { loading: submitting, run: register } = useRequest(fakeRegister, {
    manual: true,
    onSuccess: (data, params) => {
      if (data.status === "ok") {
        message.success("注册成功！");
        history.push({
          pathname: "/user/register-result",
          state: {
            account: params.email,
          },
        });
      }
    },
  });

  const onFinish = (values) => {
    register(values);
  };

  const changePrefix = (value) => {
    setPrefix(value);
  };

  return (
    <div className={styles.main}>
      <h3>注册</h3>
      <Form form={form} name="UserRegister" onFinish={onFinish}>
        <FormItem
          name="confirm"
          rules={[
            {
              required: true,
              message: "确认密码",
            },
            {
              validator: checkConfirm,
            },
          ]}
        >
          <Input size="large" type="password" placeholder="确认密码" />
          <Button
            size="large"
            loading={submitting}
            className={styles.submit}
            type="primary"
            htmlType="submit"
          >
            <span>注册</span>
          </Button>
          <Link className={styles.login} to="/user/login">
            <span>使用已有账户登录</span>
          </Link>
        </FormItem>
      </Form>
    </div>
  );
};

export default Register;
