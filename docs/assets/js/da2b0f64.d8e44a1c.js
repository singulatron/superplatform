"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9886],{65831:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>R,contentTitle:()=>k,default:()=>_,frontMatter:()=>b,metadata:()=>N,toc:()=>A});var i=a(74848),n=a(28453),l=a(91366),r=a.n(l),t=(a(6050),a(57742)),o=a.n(t),d=(a(67792),a(27362)),c=a.n(d),m=a(36683),p=a.n(m),h=a(81124),x=a.n(h),j=a(60674),u=a.n(j),g=a(23397),f=a.n(g),y=(a(26651),a(51107)),v=(a(77675),a(19365));const b={id:"add-permission-to-role",title:"Add Permission to Role",description:"Adds a specific permission to a role identified by roleId.",sidebar_label:"Add Permission to Role",hide_title:!0,hide_table_of_contents:!0,api:"eJzFVcGO2zYQ/RViTi2glZzdBih0qou2gZEAMbLeQ7FrIFxxJDGhSGZI2XUF/Xsxkm3ZXqdAcmgvtiQOh++9eTPsQGEoSPuonYUc5koFIUXwWOhSF8IjNToE7ayITkhBzqDQCm3UpUYlnnfDp4VKn+yT/YBfWk0YRKxRfGwD0k3YFPmUJJch6Mp+PMmbQgLOI0lGsFCQg1RqeVxeuQ/OICTgJckGI1KA/LG7gM0xYvEbJKD51ctYQwJWNgg5jAAhARrhKcgjtZhAKGpsJOQdxJ3nyBBJ2wr6Prk8YAL01WMmSt942JqDg3c2YOD129mM/84BvH8LCRTORrSRV6X3RheDaNmnwCHdyyPc8ycsIvR93yfw0+zVy7QPVraxdqT/RvUtB3jimkU9IkYiR1d1vIrk9TWCCxuRrDTiHmmDJH4fcv43kPoEAhYt6bgbzPUrSkKat7GG/HHNBYqyYt/BQ0AaEOoCYZ1Ag7F2bFrfxsGkvAWyg/Uz9l7WjQ7ss8kiWXdqlx4YANMezd2S4SyZcYU0tQsxf/3z7d0rYCQHoPdMfOR6CvdS1tXOo3jahzyBKJ0xbjt2Lve5LFBIq0R0n9EKWYymFSW5ZujigXDYMxbvXKWtQKu80zamh0aoUSqkqRXme08NpYKj4tLrt7gbaqBt6RgsV1cWQ3WxkZppB2kw/BK0rVojIzmbFq45yb1ciPvWe0cs+KhUHaPPsyy0HskbGUtHTSp1Bn1yOShaO9BVuEHjvAhoyhtWGJWYL4T0PqTiT9eS8OQqkk0jn3ni2ZvatQHFm+UqFasaxR+a8FkGFKWjQSjeXSFLYnSBNiBzOmB+s3wnNnfp7AxxyLNsu92mlW1TR1W23xcyWXlzc5fO0jo2hjlEpCa8Lw++mwhvZVUhpdplQ0jGWutoOOR+EhASYGuNCszSW07pXYiNtKcglRLLs3G/n7xnCnZTQ/5Pl8XeTBH/ipk3UlumM6ja7fvvEQ7ZeA6PJPLjLTBl489nU3udAHuBM3QdV/eBTN/z5y8tEs+GdQIbSZpNMV5DOvCzgryUJuC/iPXDnq/6UUy31VUuh26xOy6cNC2/QQKfcTfdZsMV9T2HX95k3wHhTLN+3SeHAcCKjCHzokAfTza/mN396fRcPqwgAbmfYNO44GTJ4YGzXwV2OW5GCPzLKl3d0nXjMOr7Y/y49NUdxxk3RrNO677v/wGGxDCf",sidebar_class_name:"put api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},k=void 0,N={id:"singulatron/add-permission-to-role",title:"Add Permission to Role",description:"Adds a specific permission to a role identified by roleId.",source:"@site/docs/singulatron/add-permission-to-role.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/add-permission-to-role",permalink:"/docs/singulatron/add-permission-to-role",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"add-permission-to-role",title:"Add Permission to Role",description:"Adds a specific permission to a role identified by roleId.",sidebar_label:"Add Permission to Role",hide_title:!0,hide_table_of_contents:!0,api:"eJzFVcGO2zYQ/RViTi2glZzdBih0qou2gZEAMbLeQ7FrIFxxJDGhSGZI2XUF/Xsxkm3ZXqdAcmgvtiQOh++9eTPsQGEoSPuonYUc5koFIUXwWOhSF8IjNToE7ayITkhBzqDQCm3UpUYlnnfDp4VKn+yT/YBfWk0YRKxRfGwD0k3YFPmUJJch6Mp+PMmbQgLOI0lGsFCQg1RqeVxeuQ/OICTgJckGI1KA/LG7gM0xYvEbJKD51ctYQwJWNgg5jAAhARrhKcgjtZhAKGpsJOQdxJ3nyBBJ2wr6Prk8YAL01WMmSt942JqDg3c2YOD129mM/84BvH8LCRTORrSRV6X3RheDaNmnwCHdyyPc8ycsIvR93yfw0+zVy7QPVraxdqT/RvUtB3jimkU9IkYiR1d1vIrk9TWCCxuRrDTiHmmDJH4fcv43kPoEAhYt6bgbzPUrSkKat7GG/HHNBYqyYt/BQ0AaEOoCYZ1Ag7F2bFrfxsGkvAWyg/Uz9l7WjQ7ss8kiWXdqlx4YANMezd2S4SyZcYU0tQsxf/3z7d0rYCQHoPdMfOR6CvdS1tXOo3jahzyBKJ0xbjt2Lve5LFBIq0R0n9EKWYymFSW5ZujigXDYMxbvXKWtQKu80zamh0aoUSqkqRXme08NpYKj4tLrt7gbaqBt6RgsV1cWQ3WxkZppB2kw/BK0rVojIzmbFq45yb1ciPvWe0cs+KhUHaPPsyy0HskbGUtHTSp1Bn1yOShaO9BVuEHjvAhoyhtWGJWYL4T0PqTiT9eS8OQqkk0jn3ni2ZvatQHFm+UqFasaxR+a8FkGFKWjQSjeXSFLYnSBNiBzOmB+s3wnNnfp7AxxyLNsu92mlW1TR1W23xcyWXlzc5fO0jo2hjlEpCa8Lw++mwhvZVUhpdplQ0jGWutoOOR+EhASYGuNCszSW07pXYiNtKcglRLLs3G/n7xnCnZTQ/5Pl8XeTBH/ipk3UlumM6ja7fvvEQ7ZeA6PJPLjLTBl489nU3udAHuBM3QdV/eBTN/z5y8tEs+GdQIbSZpNMV5DOvCzgryUJuC/iPXDnq/6UUy31VUuh26xOy6cNC2/QQKfcTfdZsMV9T2HX95k3wHhTLN+3SeHAcCKjCHzokAfTza/mN396fRcPqwgAbmfYNO44GTJ4YGzXwV2OW5GCPzLKl3d0nXjMOr7Y/y49NUdxxk3RrNO677v/wGGxDCf",sidebar_class_name:"put api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Delete a Role",permalink:"/docs/singulatron/delete-role"},next:{title:"Get Permissions by Role",permalink:"/docs/singulatron/get-permissions-by-role"}},R={},A=[];function Y(e){const s={code:"code",p:"p",...(0,n.R)(),...e.components},{Details:a}=s;return a||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(y.default,{as:"h1",className:"openapi__heading",children:"Add Permission to Role"}),"\n",(0,i.jsx)(o(),{method:"put",path:"/user-svc/role/{roleId}/permission/{permissionId}"}),"\n",(0,i.jsx)(s.p,{children:"Adds a specific permission to a role identified by roleId."}),"\n",(0,i.jsxs)(s.p,{children:["Requires the ",(0,i.jsx)(s.code,{children:"user-svc:permission:assign"})," permission."]}),"\n",(0,i.jsx)(y.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsxs)(a,{style:{marginBottom:"1rem"},className:"openapi-markdown__details","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},children:(0,i.jsx)("h3",{className:"openapi-markdown__details-summary-header-params",children:(0,i.jsx)(s.p,{children:"Path Parameters"})})}),(0,i.jsx)("div",{children:(0,i.jsxs)("ul",{children:[(0,i.jsx)(p(),{className:"paramsItem",param:{description:"Role ID",in:"path",name:"roleId",required:!0,schema:{type:"string"}}}),(0,i.jsx)(p(),{className:"paramsItem",param:{description:"Permission ID",in:"path",name:"permissionId",required:!0,schema:{type:"string"}}})]})})]}),"\n",(0,i.jsx)("div",{children:(0,i.jsx)("div",{children:(0,i.jsxs)(r(),{label:void 0,id:void 0,children:[(0,i.jsxs)(v.default,{label:"200",value:"200",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"OK"})}),(0,i.jsx)("div",{children:(0,i.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,i.jsxs)(f(),{className:"openapi-tabs__schema",children:[(0,i.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"object"})})})]})}),(0,i.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,i.jsx)(x(),{responseExample:"{}",language:"json"})})]})})})})]}),(0,i.jsxs)(v.default,{label:"401",value:"401",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Unauthorized"})}),(0,i.jsx)("div",{children:(0,i.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,i.jsxs)(f(),{className:"openapi-tabs__schema",children:[(0,i.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(u(),{collapsible:!1,name:"error",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})}),(0,i.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,i.jsx)(x(),{responseExample:'{\n  "error": "string"\n}',language:"json"})})]})})})})]}),(0,i.jsxs)(v.default,{label:"500",value:"500",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Internal Server Error"})}),(0,i.jsx)("div",{children:(0,i.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,i.jsxs)(f(),{className:"openapi-tabs__schema",children:[(0,i.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(u(),{collapsible:!1,name:"error",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})}),(0,i.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,i.jsx)(x(),{responseExample:'{\n  "error": "string"\n}',language:"json"})})]})})})})]})]})})})]})}function _(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,i.jsx)(s,{...e,children:(0,i.jsx)(Y,{...e})}):Y(e)}}}]);