import React, { useState } from "react";
import { history } from "umi";
import { DownOutlined } from "@ant-design/icons";
import { Avatar, Button, Card, Col, Dropdown, Input, InputNumber, List, Menu, Modal, message, Drawer} from "antd";
import ProDescriptions from '@ant-design/pro-descriptions';
import { PageContainer } from "@ant-design/pro-layout";
import useRequest from '@ahooksjs/use-request';
import {
  removeCart,
  queryCartList,
  deleteCartByItemId,
  updateCart,
  add2Cart,
} from "./service";
import styles from "./style.less";
import { sumBy, map } from "lodash";


export const BasicList = () => {
  const [curId, setcurId] = useState("");
  const [curName, setcurName] = useState("");
  const [curDes, setcurDes] = useState("");
  const [curPrice, setcurPrice] = useState("");
  const [curPic, setcurPic] = useState("");
  const [showDetail, setShowDetail] = useState(false);
  const [idnum, setIdnum] = useState(0);
  const [itemQuantity, setItemQuantity] = useState(0);
  const [methodType, setMethodType] = useState("");
  const { data, loading, mutate } = useRequest(queryCartList);
  
  const { run: postRun } = useRequest(
    (method, params) => {
      setIdnum(params.itemId);
      if (method === "remove") {
        setMethodType("remove");
        return deleteCartByItemId(params);
      }
      if (method === "update") {
        setMethodType("update");
        setItemQuantity(params.quantity);
        return updateCart(params);
      }
      if (method === "clear") {
        setMethodType("clear");
        return removeCart(params);
      }
    },
    {
      manual: true,
      onSuccess: (result) => {
        if (methodType == "remove") {
          const temp = data.cart.itemList.filter((item) => item.itemId !== idnum);
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
        } else if (methodType == "update") {
          const upitem = map(data.cart.itemList, (item) => {
            if (item.itemId === idnum) {
              item.quantity = itemQuantity;
            }
            return item;
          });
          const sumPrice = sumBy(upitem, (item) => item.price * item.quantity);
          const upData = {
            cart: {
              itemList: upitem,
              totalPrice: sumPrice,
            },
            msg: data.msg,
            status: data.status,
          }
          mutate(upData);
        } else {
          mutate({});
        }
      },
    }
  );
  const cart = data?.cart || [];
  const list = cart?.itemList || [];
  const price = cart?.totalPrice || 0; 
  const paginationProps = {
    showSizeChanger: true,
    showQuickJumper: true,
    total: list.length,
  };

  const deleteItem = (id) => {
    postRun("remove", {
      itemId: id,
    });
  };

  const updateItem = (id, value) => {
    postRun("update", {
      itemId: id,
      quantity: value,
    })
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

  function doClear() {
    postRun("clear", {});
  }

  const extraContent = (
    <div className={styles.extraContent}>
      <div style={{fontSize: 16, fontWeight: 600, float: 'left'}}> Total Price: {price} </div>
      <Button style={{ marginLeft: 20}} onClick={doClear}>
        Clear Cart
      </Button>
      <Button type="primary" style={{ marginLeft: 20}} onClick={handleClick}>
        Check Out
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

  const addCart = async (id) => {
    const res = await add2Cart({itemid: id, quantity: 1});
    if (res.status === "success") {
      message.success(res.msg);
      return;
    } else {
      message.error(res.msg);
    } 
  };

  const cardList = list && (
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
              onClick={() => {
                setcurId(item.itemId);
                setcurName(item.itemName);
                setcurPrice(item.price);
                setcurDes(item.description);
                setcurPic(item.picUrl);
                setShowDetail(true);
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
            title={
              <a onClick={() => {
                setcurId(item.itemId);
                setcurName(item.itemName);
                setcurPrice(item.price);
                setcurDes(item.description);
                setcurPic(item.picUrl);
                setShowDetail(true);
              }}>{item.itemName}</a>
            }
            description={item.description}
          />
          <InputNumber min={1} max={100} defaultValue={item.quantity} onChange = {(value) => updateItem(item.itemId, value)}/>
        </List.Item>
      )}
    />
  );

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
            {cardList}
          </Card>
        </div>
        <Drawer
          width={800}
          visible={showDetail}
          onClose={() => {
            setcurId("");
            setcurName("");
            setcurPrice("");
            setcurDes("");
            setcurPic("");
            setShowDetail(false);
         }}
          closable={false}
        >
          {curName && (
            <ProDescriptions
              column={1}
              title={curName}
              // request={async () => ({
              //   data: cItem || {},
              // })}
              params={{
                id: curName,
              }}
            >
              <ProDescriptions.Item>
                <p className={styles.pPic}>
                  <img
                    className={styles.drawPic}
                    alt={curName}
                    src={curPic}
                  />
                </p>
              </ProDescriptions.Item>
              <ProDescriptions.Item dataIndex="price" label="Price" valueType="price">
                {curPrice}
              </ProDescriptions.Item>
              <ProDescriptions.Item dataIndex="description" label="Description" valueType="textarea">
                {curDes}
              </ProDescriptions.Item>
              <ProDescriptions.Item>
                <Button shape="round" className={styles.addcartbtn} onClick = {() => addCart(curId)}>
                  Add to Cart
                </Button>
              </ProDescriptions.Item>
            </ProDescriptions>
          )}
        </Drawer>
      </PageContainer>
      {/* <OperationModal
        done={done}
        visible={visible}
        current={current}
        onDone={handleDone}
        onSubmit={handleSubmit}
      /> */}
    </div>
  );
};
export default BasicList;
