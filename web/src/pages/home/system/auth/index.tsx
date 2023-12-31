import { useEffect, useRef, useState } from 'react'
import type { ColumnType, ColumnGroupType } from 'antd/es/table'
import { Badge, Button, Form, Modal, message } from 'antd'
import { SearchForm, DataTable } from '@/components/Data'
import RouteBreadcrumb from '@/components/PromLayout/RouteBreadcrumb'
import { DataOption } from '@/components/Data'
import { HeightLine, PaddingLine } from '@/components/HeightLine'
import { DataOptionItem } from '@/components/Data/DataOption/DataOption'
import Detail from './child/Detail'
import EditModal from './child/EditModal'
import authOptions from './options'
import authApi from '@/apis/home/system/auth'
import { Status, StatusMap } from '@/apis/types'
import { ExclamationCircleFilled } from '@ant-design/icons'
import { ApiAuthListItem, ApiAuthListReq } from '@/apis/home/system/auth/types'
import { ActionKey } from '@/apis/data'

const { confirm } = Modal
const { authApiList, authApiDelete } = authApi
const { searchItems, operationItems } = authOptions()

const defaultPadding = 12

let timer: any

/**
 * 角色管理
 */
const Auth: React.FC = () => {
    const oprationRef = useRef<HTMLDivElement>(null)
    const [queryForm] = Form.useForm()

    const [dataSource, setDataSource] = useState<ApiAuthListItem[]>([])

    const [loading, setLoading] = useState<boolean>(false)
    const [total, setTotal] = useState<number>(0)
    const [refresh, setRefresh] = useState<boolean>(false)
    const [openDetail, setOpenDetail] = useState<boolean>(false)
    const [openEdit, setOpenEdit] = useState<boolean>(false)
    const [editId, setEditId] = useState<number | undefined>()
    const [authId, setRoleId] = useState<number>(0)
    const [search, setSearch] = useState<ApiAuthListReq>({
        page: {
            curr: 1,
            size: 10
        },
        keyword: ''
    })
    const [tableSelectedRows, setTableSelectedRows] = useState<
        ApiAuthListItem[]
    >([])

    const columns: (
        | ColumnGroupType<ApiAuthListItem>
        | ColumnType<ApiAuthListItem>
    )[] = [
        {
            title: '接口名称',
            dataIndex: 'name',
            key: 'name',
            width: 220
        },
        {
            title: '接口状态',
            dataIndex: 'status',
            key: 'status',
            width: 100,
            render: (status: Status) => {
                return (
                    <Badge
                        color={StatusMap[status].color}
                        text={StatusMap[status].text}
                    />
                )
            }
        },
        {
            // TODO 两行后省略
            title: '备注',
            dataIndex: 'remark',
            key: 'remark',
            // width: 200,
            ellipsis: true
        }
    ]

    const handlerCloseEdit = () => {
        setOpenEdit(false)
    }

    const handleEditOnOk = () => {
        handlerCloseEdit()
        handlerRefresh()
    }

    const handlerCloseDetail = () => {
        setOpenDetail(false)
    }

    // 获取数据
    const handlerGetData = () => {
        setLoading(true)
        authApiList({ ...search })
            .then((res) => {
                setDataSource(res.list)
                setTotal(res.page.total)
                console.log('res', res)
            })
            .finally(() => {
                setLoading(false)
            })
    }

    useEffect(() => {
        handlerGetData()
    }, [refresh, search])

    // 刷新
    const handlerRefresh = () => {
        setRefresh((prev) => !prev)
    }

    // 分页变化
    const handlerTablePageChange = (page: number, pageSize?: number) => {
        console.log(page, pageSize)
        setSearch({
            ...search,
            page: {
                curr: page,
                size: pageSize || 10
            }
        })
    }

    // 可以批量操作的数据
    const handlerBatchData = (
        selectedRowKeys: React.Key[],
        selectedRows: ApiAuthListItem[]
    ) => {
        console.log(selectedRowKeys, selectedRows)
        setTableSelectedRows(selectedRows)
    }

    // 处理表格操作栏的点击事件
    const handlerTableAction = (key: ActionKey, record: ApiAuthListItem) => {
        console.log(key, record)
        switch (key) {
            case ActionKey.DETAIL:
                // handlerOpenDetail()
                setOpenDetail(true)
                setRoleId(record.id)

                break
            case ActionKey.EDIT:
                setOpenEdit(true)
                setEditId(record.id)
                break
            case ActionKey.DELETE:
                confirm({
                    title: `请确认是否删除该用户?`,
                    icon: <ExclamationCircleFilled />,
                    content: '用什么影响。。。。',
                    onOk() {
                        authApiDelete({
                            id: record.id
                        }).then(() => {
                            message.success('删除成功')
                            handlerRefresh()
                        })
                    },
                    onCancel() {
                        message.info('取消操作')
                    }
                })
                break
        }
    }

    // 处理搜索表单的值变化
    const handlerSearFormValuesChange = (
        changedValues: any,
        allValues: any
    ) => {
        timer && clearTimeout(timer)
        timer = setTimeout(() => {
            setSearch({
                ...search,
                ...changedValues
            })
            console.log(changedValues, allValues)
        }, 500)
    }

    const leftOptions: DataOptionItem[] = [
        {
            key: ActionKey.BATCH_IMPORT,
            label: (
                <Button type="primary" loading={loading}>
                    批量导入
                </Button>
            )
        }
    ]

    const rightOptions: DataOptionItem[] = [
        {
            key: ActionKey.REFRESH,
            label: (
                <Button
                    type="primary"
                    loading={loading}
                    onClick={handlerRefresh}
                >
                    刷新
                </Button>
            )
        }
    ]
    //操作栏按钮
    const handleOptionClick = (val: ActionKey) => {
        console.log('val----', val)
        switch (val) {
            case ActionKey.ADD:
                setOpenEdit(true)
                setEditId(undefined)
                break
            case ActionKey.RESET:
                setSearch({
                    keyword: '',
                    page: {
                        curr: 1,
                        size: 10
                    }
                })
                break
        }
    }

    return (
        <div className="bodyContent">
            <Detail
                open={openDetail}
                onClose={handlerCloseDetail}
                authId={authId}
            />
            <EditModal
                open={openEdit}
                onClose={handlerCloseEdit}
                id={editId}
                onOk={handleEditOnOk}
            />
            <div ref={oprationRef}>
                <RouteBreadcrumb />
                <HeightLine />
                <SearchForm
                    form={queryForm}
                    items={searchItems}
                    formProps={{
                        onValuesChange: handlerSearFormValuesChange
                    }}
                />
                <HeightLine />
                <DataOption
                    queryForm={queryForm}
                    rightOptions={rightOptions}
                    leftOptions={leftOptions}
                    action={handleOptionClick}
                />
                <PaddingLine
                    padding={defaultPadding}
                    height={1}
                    borderRadius={4}
                />
            </div>
            <DataTable
                dataSource={dataSource}
                columns={columns}
                operationRef={oprationRef}
                total={total}
                loading={loading}
                operationItems={operationItems}
                pageOnChange={handlerTablePageChange}
                rowSelection={{
                    onChange: handlerBatchData,
                    selectedRowKeys: tableSelectedRows.map((item) => item.id)
                }}
                action={handlerTableAction}
            />
        </div>
    )
}

export default Auth
