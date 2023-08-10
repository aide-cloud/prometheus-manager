import type {ListReponse, Query, Reponse} from "@/apis/type";
import type {GroupItem, Status} from "@/apis/prom/prom";

// 规则组列表请求参数
export type  ListGroupRequest = {
    query?: Query
    group?: GroupItem
}

// 规则组列表响应参数
export type ListGroupReply = ListReponse & {
    groups: GroupItem[]
}

// 规则组详情响应参数
export type GroupDetailReply = Reponse & {
    group: GroupItem
}

// 规则组创建参数
export type GroupCreateItem = {
    name: string,
    remark: string,
    categoriesIds: number[],
}

// 规则组更新参数
export type GroupUpdateItem = {
    name: string,
    remark: string,
    status: Status,
    categoriesIds: number[],
}