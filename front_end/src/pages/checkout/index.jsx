import React, { useRef, useState } from "react";
import {
  Card,
  Result,
  Button,
  Descriptions,
  Divider,
  Alert,
  Statistic,
  List,
  Avatar,
} from "antd";
import { PageContainer } from "@ant-design/pro-layout";
import ProForm, {
  ProFormSelect,
  ProFormText,
  StepsForm,
  ProFormDatePicker,
} from "@ant-design/pro-form";
import { history } from 'umi';
import useRequest from '@ahooksjs/use-request';
import { queryCartList, doCheckout } from './service';
import styles from "./style.less";



const StepDescriptions = ({ stepData, bordered }) => {
  const { shipAddress, payAccount, receiverAccount, receiverName } = stepData;
  return (
    <Descriptions column={1} bordered={bordered}>
      <Descriptions.Item label="Shipping address">
        {shipAddress}
      </Descriptions.Item>
      <Descriptions.Item label="Order amount">
        <Statistic
          suffix={
            <span
              style={{
                fontSize: 14,
              }}
            >
              $
            </span>
          }
          value={amount}
          precision={2}
        />
      </Descriptions.Item>
    </Descriptions>
  );
};



const StepResult = (props) => {
  const jump2main = () => {
    const { match } = props;
    // const url = match.url === '/' ? '' : match.url;
    history.push(`/welcome/Items`);
  }
  
  const jump2history = () => {
    const { match } = props;
    // const url = match.url === '/' ? '' : match.url;
    history.push(`/history`);
  }

  return (
    <Result
      status="success"
      title="Purchase success"
      subTitle="Estimated dilivery: Two days after purchase"
      extra={
        <>
          <Button type="primary" onClick={jump2main}>
            Continue shopping
          </Button>
          <Button onClick={jump2history}>
            See order history
          </Button>
        </>
      }
      className={styles.result}
    ></Result>
  );
};

const StepForm = () => {
  // const [stepData, setStepData] = useState({
  //   shipAddress: "",
  //   payAccount: "ant-design@alipay.com",
  //   receiverAccount: "test@example.com",
  //   receiverName: "Alex",
  // });
  const { data, loading, mutate } = useRequest(queryCartList);
  const amount = data?.cart.totalPrice || "";
  const list = data?.cart.itemList || [];
  const [stepData, setStepData] = useState('');
  const [current, setCurrent] = useState(0);
  const formRef = useRef();
  const states_hash = require("./states_hash.json");

  return (
    <PageContainer
      title="Place Your Order"
      content="Following the steps to finish your purchase."
    >
      <Card bordered={false}>
        <StepsForm
          current={current}
          onCurrentChange={setCurrent}
          submitter={{
            render: (props, dom) => {
              if (props.step === 3) {
                // console.log(props);
                
                return null;
              }

              return dom;
            },
          }}
        >

          <StepsForm.StepForm
            // onFinish={async (values) => {
            //   setStepData(values);
            //   return true;
            // }}
            title="Check your order"
          >
            <div className={styles.result}>
              <Alert
                closable
                showIcon
                message="Once you place your order, you are not able to edit it"
                style={{
                  marginBottom: 24,
                }}
              />
              <List
                size="large"
                rowKey="id"
                dataSource={list}
                renderItem={(item) => (
                  <List.Item>
                    <List.Item.Meta
                      avatar={
                        <Avatar src={item.picUrl} shape="square" size="large" />
                      }
                      title={<a href={item.href}>{item.itemName}</a>}
                      description={item.description}
                    />
                    <div>Qty:{item.quantity}</div>
                  </List.Item>
                )}
              />
              {/* <StepDescriptions stepData={stepData} bordered /> */}
              <Divider
                style={{
                  margin: "24px 0",
                }}
              />
            </div>
          </StepsForm.StepForm>

          <StepsForm.StepForm
            formRef={formRef}
            title="Payment method"
            // initialValues={stepData}
            // onFinish={async (values) => {
            //   setStepData(values);
            //   return true;
            // }}
          >
            <ProFormText
              label="Card number"
              width="md"
              name="card-number"
              rules={[
                {
                  required: true,
                  message: "Please enter card number",
                },
              ]}
            />
            <ProFormText
              label="Name"
              width="md"
              name="name"
              rules={[
                {
                  required: true,
                  message: "Please input your name",
                },
              ]}
            />
            <ProForm.Group title="Expiration date" size={8}>
              <ProFormDatePicker.Month name="month" />
              {/* <ProFormDatePicker.Year name="year" /> */}
            </ProForm.Group>
            <ProFormText
              label="CVV"
              width="md"
              name="cvv"
              rules={[
                {
                  required: true,
                  message: "Please input CVV",
                },
              ]}
            />
          </StepsForm.StepForm>

          <StepsForm.StepForm
            title="Shipping address"
            onFinish={async (values) => {
              const params = {
                address: values.address,
                phone: values.phone,
                name: values.fullname,
              };
              console.log(params);
              const res = await doCheckout(params);
              console.log(res);
              if (res.status === "success") {
                return true;
              } else {
                return false;
              }
            }}
          >
            <ProFormSelect
              label="Country"
              width="md"
              name="country"
              rules={[
                {
                  required: true,
                  message: "Please choose your country",
                },
              ]}
              valueEnum={{
                "united-states": "United States",
              }}
            />
            <ProFormText
              label="Full name"
              name="fullname"
              rules={[
                {
                  required: true,
                  message: "Please input your name",
                },
              ]}
            />
            <ProFormText
              label="Phone number"
              name="phone"
              rules={[
                {
                  required: true,
                  message: "Please input your number",
                },
                {
                  pattern: /^\d{10}$/,
                  message: "Phone number invalid",
                },
              ]}
            />
            <ProFormText
              label="Address"
              name="address"
              rules={[
                {
                  required: true,
                  message: "Please input your address",
                },
              ]}
              placeholder="Street address or P.O. Box"
            />
            <ProFormText label="City" name="city" />
            <ProForm.Group size={8}>
              <ProFormSelect
                label="State"
                name="state"
                rules={[
                  {
                    required: true,
                    message: "Please choose your state",
                  },
                ]}
                valueEnum={states_hash}
              />
              <ProFormText
                label="ZIP Code"
                name="zipcode"
                rules={[
                  {
                    required: true,
                    message: "Please input zip code",
                  },
                  {
                    pattern: /^\d{5}$/,
                    message: "ZIP code invalid",
                  },
                ]}
              />
            </ProForm.Group>
          </StepsForm.StepForm>

          <StepsForm.StepForm title="Success">
            <StepResult
              onFinish={async () => {
                setCurrent(0);
                formRef.current?.resetFields();
              }}
            >
            </StepResult>
          </StepsForm.StepForm>
        </StepsForm>
        <Divider
          style={{
            margin: "40px 0 24px",
          }}
        />
        <div className={styles.desc}>
          <p>
            By placing your order, you agree to Gator Amazon's privacy notice
            and conditions of use.
          </p>
        </div>
      </Card>
    </PageContainer>
  );
};

export default StepForm;
