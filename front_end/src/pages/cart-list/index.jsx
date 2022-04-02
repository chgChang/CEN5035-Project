import React, { useState } from "react";
import { history } from "umi";
import { DownOutlined } from "@ant-design/icons";
import {
  Avatar,
  Button,
  Card,
  Col,
  Dropdown,
  Input,
  InputNumber,
  List,
  Menu,
  Modal,
  Progress,
  Radio,
  Row,
} from "antd";
import { PageContainer } from "@ant-design/pro-layout";
import useRequest from '@ahooksjs/use-request';
import OperationModal from "./components/OperationModal";
import {
  addFakeList,
  queryFakeList,
  removeFakeList,
  updateFakeList,
  queryCartList,
  deleteCartByItemId,
  updateCart,
} from "./service";
import styles from "./style.less";
import { sumBy } from "lodash";
const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;
const { Search } = Input;

export const BasicList = () => {
  const [done, setDone] = useState(false);
  const [visible, setVisible] = useState(false);
  const [current, setCurrent] = useState(undefined);
  const [idnum, setIdnum] = useState(0);
  const { data, loading, mutate } = useRequest(queryCartList);
  console.log(data);
  
  const { run: postRun } = useRequest(
    (method, params) => {
      if (method === "remove") {
        setIdnum(params.itemId);
        console.log("remove" + idnum);
        return deleteCartByItemId(params);
      }
      if (method === "update") {
        return updateFakeList(params);
      }
    },
    {
      manual: true,
      onSuccess: (result) => {
        console.log(data);
        const temp = data.cart.itemList.filter((item) => item.itemId !== idnum);
        const delitem = {
          cart: data.cart.itemList.filter((item) => item.itemId === idnum),
        };
        const sumPrice = sumBy(temp, (item) => item.price * item.quantity);
        const delData = {
          cart: {
            itemList: temp,
            totalPrice: sumPrice,
          },
          msg: data.msg,
          status: data.status,
        }
        mutate(delData);
      },
    }
  );
  const list = data?.cart.itemList || [];
  const price = data?.cart.totalPrice || 0; 
  const paginationProps = {
    showSizeChanger: true,
    showQuickJumper: true,
    total: list.length,
  };

  const showEditModal = (item) => {
    setVisible(true);
    setCurrent(item);
  };

  const deleteItem = (id) => {
    // console.log(id);
    postRun("remove", {
      itemId: id,
    });
  };

  const doDelete = (item) => {
    Modal.confirm({
      title: "Delete item",
      content: "Are you sure to delete this item？",
      okText: "Ok",
      cancelText: "Cancel",
      onOk: () => deleteItem(item.itemId),
    });
  };

  function handleClick() {
    history.push("/checkout");
  }

  const extraContent = (
    <div className={styles.extraContent}>
      <div style={{fontSize: 16, fontWeight: 600, float: 'left'}}> Total Price: {price} </div>
      <Button type="primary" style={{ marginLeft: 20}} onClick={handleClick}>
        Proceed to checkout
      </Button>
    </div>
  );

  const MoreBtn = ({ item }) => (
    <Dropdown
      overlay={
        <Menu onClick={() => doDelete(item)}>
          {/* <Menu.Item key="edit">xxx</Menu.Item> */}
          <Menu.Item key="delete">Delete</Menu.Item>
        </Menu>
      }
    >
      <a>
        More <DownOutlined />
      </a>
    </Dropdown>
  );

  const handleDone = () => {
    setDone(false);
    setVisible(false);
    setCurrent({});
  };

  const handleSubmit = (values) => {
    setDone(true);
    const method = values?.id ? "update" : "add";
    postRun(method, values);
  };

  const updateItem = (id, value) => {
    console.log("id" + id);
    console.log("quantity" + value);
    const info = {
      itemId: id,
      quantity: value,
    }
    updateCart(info);
  };

  return (
    <div>
      <PageContainer>
        <div className={styles.standardList}>
          <Card
            className={styles.listCard}
            bordered={false}
            title="Items"
            style={{
              marginTop: 24,
            }}
            bodyStyle={{
              padding: "0 32px 40px 32px",
            }}
            extra={extraContent}
          >
            <List
              size="large"
              rowKey="id"
              loading={loading}
              pagination={paginationProps}
              dataSource={list}
              renderItem={(item) => (
                <List.Item
                  actions={[
                    <a
                      key="edit"
                      onClick={(e) => {
                        e.preventDefault();
                        //showEditModal(item);
                      }}
                    >
                      Details
                    </a>,
                    <MoreBtn key="more" item={item} />,
                  ]}
                >
                  <List.Item.Meta
                    avatar={
                      <Avatar src={item.picUrl} shape="square" size="large" />
                    }
                    title={<a href={item.href}>{item.itemName}</a>}
                    description={item.description}
                  />
                  <InputNumber min={1} max={100} defaultValue={item.quantity} onChange = {(value) => updateItem(item.itemId, value)}/>
                </List.Item>
              )}
            />
          </Card>
        </div>
      </PageContainer>
      <OperationModal
        done={done}
        visible={visible}
        current={current}
        onDone={handleDone}
        onSubmit={handleSubmit}
      />
    </div>
  );
};
export default BasicList;
