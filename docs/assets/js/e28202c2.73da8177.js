"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4029],{62922:(e,s,i)=>{i.r(s),i.d(s,{assets:()=>w,contentTitle:()=>N,default:()=>S,frontMatter:()=>f,metadata:()=>_,toc:()=>R});var a=i(74848),r=i(28453),l=i(91366),n=i.n(l),t=(i(6050),i(57742)),d=i.n(t),o=(i(67792),i(27362)),c=i.n(o),m=i(36683),p=i.n(m),h=i(81124),j=i.n(h),u=i(60674),x=i.n(u),g=i(23397),b=i.n(g),v=(i(26651),i(51107)),y=(i(77675),i(19365));const f={id:"get-permissions-by-role",title:"Get Permissions by Role",description:"Retrieve permissions associated with a specific role ID.",sidebar_label:"Get Permissions by Role",hide_title:!0,hide_table_of_contents:!0,api:"eJylVd9v2zYQ/leIe9oARXKTFRj0tBTrAq/FEsTJgCLxw1k6S2wpkj1S9jxD//twkh0rsVNgyIstkvfzu/vutlBSKFj7qJ2FHG4psqYVKU/c6BC0s0FhCK7QGKlUax1rhSp4KvRSF4qdITX9PYUEPDI2FIkD5A/bl3YHOUhAy9FjrCEBiw1BDmJkWkICTN9bzVRCHrmlBEJRU4OQbyFuvEhqG6kihq6bi3TwzgYKInA+mcjfc6/XnyCBwtlINsorem90gfKafQ0ish358Ow8cdSDwREActSRmnAsVTAJLpdxFGSIrG0FXfI8mhPvujyOmapUPUIbiNOVpjXxI0ByrDpAd1r5PhCrv3+g7NaWeHrC+Yx4pQtS69opt7ZBxXrcCqdstb58FYHuSd4tvlIR4XCBzLg5JdHJ3S+nqjm1KzS6VH/Orv/6P3V9GdTg4N2xg3uLbawd63+pfKuD96cziMQWjRKcidVHZsdv89QlEKhoWcdNT7sPhEx82cYa8oe50CRiJYwcumJXYJgn0FCsXQk5VBR7+ooKZNJ5Z2EQy4SZ2XbgZ5eNKSFuJYeB7C0b0c2MK9DULsT8/a/nF+9A/O/Dm0kWA2vGQb7E6G7jST3uRB5BLZ0xbk2lWmz6wYMFKbSliu4bWYXFMDHUkl3Td2ufZtg38mdXaavIlt5pG9P9/KkJS+LDBLrc1b3H/dDl6PUnGrpU26XrCe9sxKIvFTWoJe2AhsJvQduqNRjZ2bRwzcj2zVTNWu8dC8wDUnWMPs+y0HpibzAuHTcp6uxoaMBta/t0S1qRcV4FMsszQZhKdTlV6H1I1RfXsvLsKsamwYUhpe1Z7dpA6urmLlV3Nak/NNMCA6ml4x4o0a5IIDG6IBv6cbKP+erms1pdpJNnEYc8y9brdVrZNnVcZTu9kGHlzdlFOknr2Jie5MRNuF7uu+2Q8BqrijjVLutFMsFaRyMiswOAkIC01oDAJD0Xk96F2KAdBXlFUd2M1tRio2TNwNHcfaLXW9bbriUi/RMzb1DbfvgJNtsddx5gzB1ZZ0M0+dN+GxNonoCUUdS2WynMPZuuk+vvLbGQeZ7ACllLPYeNqoN8l5Av0QT6QZo/3e4W6c/qsHhPJrBvdLsRzNG0coIEvtHmsJi7eZfsKSOBDI+XRUE+jtSORlc3njJXH+8gAdxx/kAwMZbsP8T6yZBeEnQIQX675BWV7Xagb9c9yQ9Pr2o8TYVBWhCad133HwjjIrw=",sidebar_class_name:"get api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},N=void 0,_={id:"singulatron/get-permissions-by-role",title:"Get Permissions by Role",description:"Retrieve permissions associated with a specific role ID.",source:"@site/docs/singulatron/get-permissions-by-role.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/get-permissions-by-role",permalink:"/docs/singulatron/get-permissions-by-role",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"get-permissions-by-role",title:"Get Permissions by Role",description:"Retrieve permissions associated with a specific role ID.",sidebar_label:"Get Permissions by Role",hide_title:!0,hide_table_of_contents:!0,api:"eJylVd9v2zYQ/leIe9oARXKTFRj0tBTrAq/FEsTJgCLxw1k6S2wpkj1S9jxD//twkh0rsVNgyIstkvfzu/vutlBSKFj7qJ2FHG4psqYVKU/c6BC0s0FhCK7QGKlUax1rhSp4KvRSF4qdITX9PYUEPDI2FIkD5A/bl3YHOUhAy9FjrCEBiw1BDmJkWkICTN9bzVRCHrmlBEJRU4OQbyFuvEhqG6kihq6bi3TwzgYKInA+mcjfc6/XnyCBwtlINsorem90gfKafQ0ish358Ow8cdSDwREActSRmnAsVTAJLpdxFGSIrG0FXfI8mhPvujyOmapUPUIbiNOVpjXxI0ByrDpAd1r5PhCrv3+g7NaWeHrC+Yx4pQtS69opt7ZBxXrcCqdstb58FYHuSd4tvlIR4XCBzLg5JdHJ3S+nqjm1KzS6VH/Orv/6P3V9GdTg4N2xg3uLbawd63+pfKuD96cziMQWjRKcidVHZsdv89QlEKhoWcdNT7sPhEx82cYa8oe50CRiJYwcumJXYJgn0FCsXQk5VBR7+ooKZNJ5Z2EQy4SZ2XbgZ5eNKSFuJYeB7C0b0c2MK9DULsT8/a/nF+9A/O/Dm0kWA2vGQb7E6G7jST3uRB5BLZ0xbk2lWmz6wYMFKbSliu4bWYXFMDHUkl3Td2ufZtg38mdXaavIlt5pG9P9/KkJS+LDBLrc1b3H/dDl6PUnGrpU26XrCe9sxKIvFTWoJe2AhsJvQduqNRjZ2bRwzcj2zVTNWu8dC8wDUnWMPs+y0HpibzAuHTcp6uxoaMBta/t0S1qRcV4FMsszQZhKdTlV6H1I1RfXsvLsKsamwYUhpe1Z7dpA6urmLlV3Nak/NNMCA6ml4x4o0a5IIDG6IBv6cbKP+erms1pdpJNnEYc8y9brdVrZNnVcZTu9kGHlzdlFOknr2Jie5MRNuF7uu+2Q8BqrijjVLutFMsFaRyMiswOAkIC01oDAJD0Xk96F2KAdBXlFUd2M1tRio2TNwNHcfaLXW9bbriUi/RMzb1DbfvgJNtsddx5gzB1ZZ0M0+dN+GxNonoCUUdS2WynMPZuuk+vvLbGQeZ7ACllLPYeNqoN8l5Av0QT6QZo/3e4W6c/qsHhPJrBvdLsRzNG0coIEvtHmsJi7eZfsKSOBDI+XRUE+jtSORlc3njJXH+8gAdxx/kAwMZbsP8T6yZBeEnQIQX675BWV7Xagb9c9yQ9Pr2o8TYVBWhCad133HwjjIrw=",sidebar_class_name:"get api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Add Permission to Role",permalink:"/docs/singulatron/add-permission-to-role"},next:{title:"Set Role Permissions",permalink:"/docs/singulatron/set-role-permissions"}},w={},R=[];function I(e){const s={p:"p",...(0,r.R)(),...e.components},{Details:i}=s;return i||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(v.default,{as:"h1",className:"openapi__heading",children:"Get Permissions by Role"}),"\n",(0,a.jsx)(d(),{method:"get",path:"/user-service/role/{roleId}/permissions"}),"\n",(0,a.jsx)(s.p,{children:"Retrieve permissions associated with a specific role ID."}),"\n",(0,a.jsx)(v.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,a.jsxs)(i,{style:{marginBottom:"1rem"},className:"openapi-markdown__details","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},children:(0,a.jsx)("h3",{className:"openapi-markdown__details-summary-header-params",children:(0,a.jsx)(s.p,{children:"Path Parameters"})})}),(0,a.jsx)("div",{children:(0,a.jsx)("ul",{children:(0,a.jsx)(p(),{className:"paramsItem",param:{description:"Role ID",in:"path",name:"roleId",required:!0,schema:{type:"integer"}}})})})]}),"\n",(0,a.jsx)("div",{children:(0,a.jsx)("div",{children:(0,a.jsxs)(n(),{label:void 0,id:void 0,children:[(0,a.jsxs)(y.default,{label:"200",value:"200",children:[(0,a.jsx)("div",{children:(0,a.jsx)(s.p,{children:"OK"})}),(0,a.jsx)("div",{children:(0,a.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,a.jsx)(y.default,{label:"application/json",value:"application/json",children:(0,a.jsxs)(b(),{className:"openapi-tabs__schema",children:[(0,a.jsx)(y.default,{label:"Schema",value:"Schema",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,a.jsx)("strong",{children:(0,a.jsx)(s.p,{children:"Schema"})})}),(0,a.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,a.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,a.jsx)(x(),{collapsible:!0,className:"schemaItem",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details",children:[(0,a.jsx)("summary",{style:{},children:(0,a.jsxs)("span",{className:"openapi-schema__container",children:[(0,a.jsx)("strong",{className:"openapi-schema__property",children:(0,a.jsx)(s.p,{children:"permissions"})}),(0,a.jsx)("span",{className:"openapi-schema__name",children:(0,a.jsx)(s.p,{children:"object[]"})})]})}),(0,a.jsxs)("div",{style:{marginLeft:"1rem"},children:[(0,a.jsx)("li",{children:(0,a.jsx)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem",paddingBottom:".5rem"},children:(0,a.jsx)(s.p,{children:"Array ["})})}),(0,a.jsx)(x(),{collapsible:!1,name:"createdAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,a.jsx)(x(),{collapsible:!1,name:"description",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,a.jsx)(x(),{collapsible:!1,name:"id",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:'eg. "user.viewer"',type:"string"}}),(0,a.jsx)(x(),{collapsible:!1,name:"name",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:'eg. "User Viewer"',type:"string"}}),(0,a.jsx)(x(),{collapsible:!1,name:"ownerId",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:"Service who owns the permission",type:"string"}}),(0,a.jsx)(x(),{collapsible:!1,name:"updatedAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,a.jsx)("li",{children:(0,a.jsx)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem"},children:(0,a.jsx)(s.p,{children:"]"})})})]})]})})})]})}),(0,a.jsx)(y.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,a.jsx)(j(),{responseExample:'{\n  "permissions": [\n    {\n      "createdAt": "string",\n      "description": "string",\n      "id": "string",\n      "name": "string",\n      "ownerId": "string",\n      "updatedAt": "string"\n    }\n  ]\n}',language:"json"})})]})})})})]}),(0,a.jsxs)(y.default,{label:"400",value:"400",children:[(0,a.jsx)("div",{children:(0,a.jsx)(s.p,{children:"Invalid JSON"})}),(0,a.jsx)("div",{children:(0,a.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,a.jsx)(y.default,{label:"application/json",value:"application/json",children:(0,a.jsx)(b(),{className:"openapi-tabs__schema",children:(0,a.jsx)(y.default,{label:"Schema",value:"Schema",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,a.jsx)("strong",{children:(0,a.jsx)(s.p,{children:"Schema"})})}),(0,a.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,a.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,a.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,a.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,a.jsxs)(y.default,{label:"401",value:"401",children:[(0,a.jsx)("div",{children:(0,a.jsx)(s.p,{children:"Unauthorized"})}),(0,a.jsx)("div",{children:(0,a.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,a.jsx)(y.default,{label:"application/json",value:"application/json",children:(0,a.jsx)(b(),{className:"openapi-tabs__schema",children:(0,a.jsx)(y.default,{label:"Schema",value:"Schema",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,a.jsx)("strong",{children:(0,a.jsx)(s.p,{children:"Schema"})})}),(0,a.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,a.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,a.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,a.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,a.jsxs)(y.default,{label:"500",value:"500",children:[(0,a.jsx)("div",{children:(0,a.jsx)(s.p,{children:"Internal Server Error"})}),(0,a.jsx)("div",{children:(0,a.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,a.jsx)(y.default,{label:"application/json",value:"application/json",children:(0,a.jsx)(b(),{className:"openapi-tabs__schema",children:(0,a.jsx)(y.default,{label:"Schema",value:"Schema",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,a.jsx)("strong",{children:(0,a.jsx)(s.p,{children:"Schema"})})}),(0,a.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,a.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,a.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,a.jsx)(s.p,{children:"string"})})})]})})})})})})]})]})})})]})}function S(e={}){const{wrapper:s}={...(0,r.R)(),...e.components};return s?(0,a.jsx)(s,{...e,children:(0,a.jsx)(I,{...e})}):I(e)}}}]);