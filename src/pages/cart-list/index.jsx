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
import { useRequest } from "umi";
import moment from "moment";
import OperationModal from "./components/OperationModal";
import {
  addFakeList,
  queryFakeList,
  removeFakeList,
  updateFakeList,
} from "./service";
import styles from "./style.less";
const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;
const { Search } = Input;

const Info = ({ title, value, bordered }) => (
  <div className={styles.headerInfo}>
    <span>{title}</span>
    <p>{value}</p>
    {bordered && <em />}
  </div>
);

const ListContent = ({ data: { owner, createdAt, percent, itemNumber } }) => (
  <div className={styles.listContent}>
    <div className={styles.listContentItem}>
      <InputNumber min={1} max={100} defaultValue={itemNumber} />
    </div>
  </div>
);

export const BasicList = () => {
  const [done, setDone] = useState(false);
  const [visible, setVisible] = useState(false);
  const [current, setCurrent] = useState(undefined);
  const {
    data: listData,
    loading,
    mutate,
  } = useRequest(() => {
    return queryFakeList({
      count: 50,
    });
  });
  const { run: postRun } = useRequest(
    (method, params) => {
      if (method === "remove") {
        return removeFakeList(params);
      }

      if (method === "update") {
        return updateFakeList(params);
      }

      return addFakeList(params);
    },
    {
      manual: true,
      onSuccess: (result) => {
        mutate(result);
      },
    }
  );
  const list = listData?.list || [];
  const paginationProps = {
    showSizeChanger: true,
    showQuickJumper: true,
    pageSize: 5,
    total: list.length,
  };

  const showEditModal = (item) => {
    setVisible(true);
    setCurrent(item);
  };

  const deleteItem = (id) => {
    postRun("remove", {
      id,
    });
  };

  const editAndDelete = (key, currentItem) => {
    if (key === "edit") showEditModal(currentItem);
    else if (key === "delete") {
      Modal.confirm({
        title: "Delete item",
        content: "Are you sure to delete this item？",
        okText: "Ok",
        cancelText: "Cancel",
        onOk: () => deleteItem(currentItem.id),
      });
    }
  };

  function handleClick() {
    history.push("/checkout");
  }

  const extraContent = (
    <div className={styles.extraContent}>
      <Button type="primary" onClick={handleClick}>
        Proceed to checkout
      </Button>
      {/* <Search className={styles.extraContentSearch} placeholder="Input item name" onSearch={() => ({})} /> */}
    </div>
  );

  const MoreBtn = ({ item }) => (
    <Dropdown
      overlay={
        <Menu onClick={({ key }) => editAndDelete(key, item)}>
          <Menu.Item key="edit">xxx</Menu.Item>
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
                        showEditModal(item);
                      }}
                    >
                      Details
                    </a>,
                    <MoreBtn key="more" item={item} />,
                  ]}
                >
                  <List.Item.Meta
                    avatar={
                      <Avatar src={item.logo} shape="square" size="large" />
                    }
                    title={<a href={item.href}>{item.title}</a>}
                    description={item.subDescription}
                  />
                  <ListContent data={item} />
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
