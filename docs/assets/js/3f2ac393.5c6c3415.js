"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9350],{67365:(e,a,s)=>{s.r(a),s.d(a,{assets:()=>_,contentTitle:()=>v,default:()=>q,frontMatter:()=>g,metadata:()=>b,toc:()=>w});var i=s(74848),r=s(28453),t=s(91366),l=s.n(t),n=(s(6050),s(57742)),d=s.n(n),c=(s(67792),s(27362)),o=s.n(c),h=(s(36683),s(81124)),m=s.n(h),p=s(60674),x=s.n(p),j=s(23397),u=s.n(j),y=(s(26651),s(51107)),f=(s(77675),s(19365));const g={id:"retrieve-details-of-a-chat-thread",title:"Retrieve details of a chat thread",description:"Fetch information about a specific chat thread by its ID",sidebar_label:"Retrieve details of a chat thread",hide_title:!0,hide_table_of_contents:!0,api:"eJylVcFy1DAM/RWPzqZbGLjkVqAwy4Ey7fa07EGxtYlLYqe2sm3Yyb8zSrJ0WzIdOs0ljixL78nSyx4sJRNdwy54yOALsSmV89sQaxSbwjy0rFClhozbOqNMiay4jIRW5Z1ynNTyM2iIdNtS4o/BdpDtwQTP5FmW2DSVM0O4xU2SPHtIpqQaZdXE0FBkR0m+xsBLO6y7hiCDxNH5AvpeHywhvyHD0IvpMf6vxGo1Yrsc8UzIXCQLGceWejGkJvg0Znx3eiqvx3GmGJYYXZVUao2hlLZtVXUqEkdHO7KgX8fyX7uJhEz2jGfoa3BzVdHAjiuaoSBmFbaKS5ru6wT0zPHQOLO0aSbCtKMsbZ2npO5KZ0o1HEhHYVVOVfBFUhxOfvrVuG3Qqxjaoqw6lYunrFnwYFKMRVLbEKcASYA5pjrNExwNGCN28t029pkqtYniLJ3rcWMA7mw6lEb8hVoYICcSrC4dVex/gc3050zH9hrez3Xc0u+wclZ9u7r4/pLGejokY4K3M/Q9tlyG6H6/rHPnEnyYZ8AUPVbqiuKOojqPMcTXZZICYpEgW4PIDmw01MRlsJBBE4bhbpBLyGAh+4tprjRI8MsHQTq/x7oZh+RBYI4ny2/DQbTQCNSHuQLQsKOYRpLiLZlrHMB7rMXlcpKEv4IhbX6slPBEqI708TWaO1WM6Z4XTYXODwMQq0FZhsJMhdMHydloKKVw2Rr2+xwTXceq78V821LsIFtvNOwwOsyF/HrTaygJLUXI1nv4RR1k8GnE/mYl6cW9aoc5eHq7vT6cODOGGn7W9/huf1xcrUBDPv1L6mDlTMQ7UXO8Gy8lDLUcJnOw7aFCX7RYiO8YU54/Hxxk5g==",sidebar_class_name:"post api-method",custom_edit_url:null},v=void 0,b={id:"singulatron/retrieve-details-of-a-chat-thread",title:"Retrieve details of a chat thread",description:"Fetch information about a specific chat thread by its ID",source:"@site/docs/singulatron/retrieve-details-of-a-chat-thread.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/retrieve-details-of-a-chat-thread",permalink:"/docs/singulatron/retrieve-details-of-a-chat-thread",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"retrieve-details-of-a-chat-thread",title:"Retrieve details of a chat thread",description:"Fetch information about a specific chat thread by its ID",sidebar_label:"Retrieve details of a chat thread",hide_title:!0,hide_table_of_contents:!0,api:"eJylVcFy1DAM/RWPzqZbGLjkVqAwy4Ey7fa07EGxtYlLYqe2sm3Yyb8zSrJ0WzIdOs0ljixL78nSyx4sJRNdwy54yOALsSmV89sQaxSbwjy0rFClhozbOqNMiay4jIRW5Z1ynNTyM2iIdNtS4o/BdpDtwQTP5FmW2DSVM0O4xU2SPHtIpqQaZdXE0FBkR0m+xsBLO6y7hiCDxNH5AvpeHywhvyHD0IvpMf6vxGo1Yrsc8UzIXCQLGceWejGkJvg0Znx3eiqvx3GmGJYYXZVUao2hlLZtVXUqEkdHO7KgX8fyX7uJhEz2jGfoa3BzVdHAjiuaoSBmFbaKS5ru6wT0zPHQOLO0aSbCtKMsbZ2npO5KZ0o1HEhHYVVOVfBFUhxOfvrVuG3Qqxjaoqw6lYunrFnwYFKMRVLbEKcASYA5pjrNExwNGCN28t029pkqtYniLJ3rcWMA7mw6lEb8hVoYICcSrC4dVex/gc3050zH9hrez3Xc0u+wclZ9u7r4/pLGejokY4K3M/Q9tlyG6H6/rHPnEnyYZ8AUPVbqiuKOojqPMcTXZZICYpEgW4PIDmw01MRlsJBBE4bhbpBLyGAh+4tprjRI8MsHQTq/x7oZh+RBYI4ny2/DQbTQCNSHuQLQsKOYRpLiLZlrHMB7rMXlcpKEv4IhbX6slPBEqI708TWaO1WM6Z4XTYXODwMQq0FZhsJMhdMHydloKKVw2Rr2+xwTXceq78V821LsIFtvNOwwOsyF/HrTaygJLUXI1nv4RR1k8GnE/mYl6cW9aoc5eHq7vT6cODOGGn7W9/huf1xcrUBDPv1L6mDlTMQ7UXO8Gy8lDLUcJnOw7aFCX7RYiO8YU54/Hxxk5g==",sidebar_class_name:"post api-method",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Retrieve messages from a chat thread",permalink:"/docs/singulatron/retrieve-messages-from-a-chat-thread"},next:{title:"Create a new chat thread",permalink:"/docs/singulatron/create-a-new-chat-thread"}},_={},w=[];function N(e){const a={p:"p",...(0,r.R)(),...e.components},{Details:s}=a;return s||function(e,a){throw new Error("Expected "+(a?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(y.default,{as:"h1",className:"openapi__heading",children:"Retrieve details of a chat thread"}),"\n",(0,i.jsx)(d(),{method:"post",path:"/chat/thread"}),"\n",(0,i.jsx)(a.p,{children:"Fetch information about a specific chat thread by its ID"}),"\n",(0,i.jsx)(y.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(o(),{className:"openapi-tabs__mime",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json-schema",children:(0,i.jsxs)(s,{style:{},className:"openapi-markdown__details mime","data-collapsed":!1,open:!0,children:[(0,i.jsxs)("summary",{style:{},className:"openapi-markdown__details-summary-mime",children:[(0,i.jsx)("h3",{className:"openapi-markdown__details-summary-header-body",children:(0,i.jsx)(a.p,{children:"Body"})}),(0,i.jsx)("strong",{className:"openapi-schema__required",children:(0,i.jsx)(a.p,{children:"required"})})]}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:"1rem",marginBottom:"1rem"},children:(0,i.jsx)(a.p,{children:"Get Thread Request"})})}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(x(),{collapsible:!1,name:"threadId",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})})}),"\n",(0,i.jsx)("div",{children:(0,i.jsx)("div",{children:(0,i.jsxs)(l(),{label:void 0,id:void 0,children:[(0,i.jsxs)(f.default,{label:"200",value:"200",children:[(0,i.jsx)("div",{children:(0,i.jsx)(a.p,{children:"Thread details successfully retrieved"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsxs)(u(),{className:"openapi-tabs__schema",children:[(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(s,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(a.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(x(),{collapsible:!0,className:"schemaItem",children:(0,i.jsxs)(s,{style:{},className:"openapi-markdown__details",children:[(0,i.jsx)("summary",{style:{},children:(0,i.jsxs)("span",{className:"openapi-schema__container",children:[(0,i.jsx)("strong",{className:"openapi-schema__property",children:(0,i.jsx)(a.p,{children:"thread"})}),(0,i.jsx)("span",{className:"openapi-schema__name",children:(0,i.jsx)(a.p,{children:"object"})})]})}),(0,i.jsxs)("div",{style:{marginLeft:"1rem"},children:[(0,i.jsx)(x(),{collapsible:!1,name:"createdAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"id",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"title",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:"Title of the thread.",type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"topicIds",required:!1,schemaName:"string[]",qualifierMessage:void 0,schema:{description:"TopicIds defines which topics the thread belongs to.\nTopics can roughly be thought of as tags for threads.",items:{type:"string"},type:"array"}}),(0,i.jsx)(x(),{collapsible:!1,name:"updatedAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"userIds",required:!1,schemaName:"string[]",qualifierMessage:void 0,schema:{description:"UserIds the ids of the users who can see this thread.",items:{type:"string"},type:"array"}})]})]})})})]})}),(0,i.jsx)(f.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,i.jsx)(m(),{responseExample:'{\n  "thread": {\n    "createdAt": "string",\n    "id": "string",\n    "title": "string",\n    "topicIds": [\n      "string"\n    ],\n    "updatedAt": "string",\n    "userIds": [\n      "string"\n    ]\n  }\n}',language:"json"})})]})})})})]}),(0,i.jsxs)(f.default,{label:"400",value:"400",children:[(0,i.jsx)("div",{children:(0,i.jsx)(a.p,{children:"Invalid JSON"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(u(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(s,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(a.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(a.p,{children:"string"})})})]})})})})})})]}),(0,i.jsxs)(f.default,{label:"401",value:"401",children:[(0,i.jsx)("div",{children:(0,i.jsx)(a.p,{children:"Unauthorized"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(u(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(s,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(a.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(a.p,{children:"string"})})})]})})})})})})]}),(0,i.jsxs)(f.default,{label:"500",value:"500",children:[(0,i.jsx)("div",{children:(0,i.jsx)(a.p,{children:"Internal Server Error"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(u(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(s,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(a.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(a.p,{children:"string"})})})]})})})})})})]})]})})})]})}function q(e={}){const{wrapper:a}={...(0,r.R)(),...e.components};return a?(0,i.jsx)(a,{...e,children:(0,i.jsx)(N,{...e})}):N(e)}}}]);