"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4029],{62922:(e,s,i)=>{i.r(s),i.d(s,{assets:()=>N,contentTitle:()=>_,default:()=>k,frontMatter:()=>b,metadata:()=>S,toc:()=>I});var a=i(74848),n=i(28453),r=i(91366),l=i.n(r),t=(i(6050),i(57742)),d=i.n(t),o=(i(67792),i(27362)),m=i.n(o),c=i(36683),p=i.n(c),h=i(81124),u=i.n(h),x=i(60674),j=i.n(x),y=i(23397),g=i.n(y),v=(i(26651),i(51107)),f=(i(77675),i(19365));const b={id:"get-permissions-by-role",title:"Get Permissions by Role",description:"Retrieve permissions associated with a specific role ID.",sidebar_label:"Get Permissions by Role",hide_title:!0,hide_table_of_contents:!0,api:"eJylVdtu20YQ/ZXFPLUATTp2AxR8qoOmhpqgNiy7QGHrYUyOyE2Wu5vZpVRF4L8XQ+pCS3KAwi8SyZ3LmbNzZtZQUihY+6idhRzuKLKmBSlP3OgQtLNBYQiu0BipVEsda4UqeCr0XBeKnSE1+T2FBJwnRgkzKSGHiuLtPsSH1Z0zBAl4ZGwoEgfIH9eHyYdgkICWV4+xhgQsNgQ5SKZJCQkwfWs1Uwl55JYSCEVNDUK+hrjyYqltpIoYum4m1sE7GyiIwcX5ufy9zHrzCRIonI1ko5yi90YXfSXZlyAm61EOz1Jn1EPAEUvyqiM14diqYBLyruIIZIisbQVd8hLNiXNdHmOmKlVP0AbidKFpSfwEkBy7DtSddn4IxOrvHzi7pSWenEg+JV7ogtSydsotbVCxHvfLqVitL19loNvZu+cvVETYf0BmXJ2y6OTbL6duc2IXaHSp/pze/PV/7vUQ1JDg3XGCB4ttrB3r71S+NcH70xVEYotGCc/E6iOz47dl6hIIVLSs46qX3QdCJr5qYw3540xkErESRQ5dsblgmCXQUKzdRs+9fMUFMum8s7AoMlFlth602WVjOUhKwT8IvWUjfplxBZrahZi///Xi8h1I7i20qVQwKGYM8JCf+5Un9bQxeQI1d8a4JZXqedVPJixIoS1VdF/JKiyGaaHm7Jq+U/sSw7aJP7tKW0W29E7bmG5nT01YEu+nz9XmznvO9x2OXn+ioUO1nbte7M5GLPproga1lB3QUPgtaFu1BiM7mxauGcW+nahp671joXhgqo7R51kWWk/sDca54yZFnR0NDLhrbV9uSQsyzqtAZn4mDFOpriYKvQ+p+se1rDy7irFp8NmQ0vasdm0gdX17n6r7mtQfmukZA6m5454o8a5IKDG6IBv6UbLFfH37WS0u0/MXiEOeZcvlMq1smzquso1fyLDy5uwyPU/r2Jhe4MRNuJlvO21f8BKrijjVLutNMuFaRyMm0z2BkIC01sDAeXohIb0LsUE7AnlNUY2WkDTIZg0dzNydtN6y/zYtEenfmHmD2vaDT7hZb3TzCFvdyBobkOS7vTYWzywBuUJxWa/lUh7YdJ18/tYSi4hnCSyQtdzlsEl1kOcS8jmaQD8o8ae7zQL9We0X7knw2ya3K+EbTStvkMBXWu0Xcjfrkq1cBMhweFUU5OPI7WhkdePpcv3xHhLAjd734pJgyfZBop+EdCjOAYL8dskrLuv1IN2u29kPR6967CbCYC0Mzbqu+w+2WC0H",sidebar_class_name:"get api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},_=void 0,S={id:"singulatron/get-permissions-by-role",title:"Get Permissions by Role",description:"Retrieve permissions associated with a specific role ID.",source:"@site/docs/singulatron/get-permissions-by-role.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/get-permissions-by-role",permalink:"/docs/singulatron/get-permissions-by-role",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"get-permissions-by-role",title:"Get Permissions by Role",description:"Retrieve permissions associated with a specific role ID.",sidebar_label:"Get Permissions by Role",hide_title:!0,hide_table_of_contents:!0,api:"eJylVdtu20YQ/ZXFPLUATTp2AxR8qoOmhpqgNiy7QGHrYUyOyE2Wu5vZpVRF4L8XQ+pCS3KAwi8SyZ3LmbNzZtZQUihY+6idhRzuKLKmBSlP3OgQtLNBYQiu0BipVEsda4UqeCr0XBeKnSE1+T2FBJwnRgkzKSGHiuLtPsSH1Z0zBAl4ZGwoEgfIH9eHyYdgkICWV4+xhgQsNgQ5SKZJCQkwfWs1Uwl55JYSCEVNDUK+hrjyYqltpIoYum4m1sE7GyiIwcX5ufy9zHrzCRIonI1ko5yi90YXfSXZlyAm61EOz1Jn1EPAEUvyqiM14diqYBLyruIIZIisbQVd8hLNiXNdHmOmKlVP0AbidKFpSfwEkBy7DtSddn4IxOrvHzi7pSWenEg+JV7ogtSydsotbVCxHvfLqVitL19loNvZu+cvVETYf0BmXJ2y6OTbL6duc2IXaHSp/pze/PV/7vUQ1JDg3XGCB4ttrB3r71S+NcH70xVEYotGCc/E6iOz47dl6hIIVLSs46qX3QdCJr5qYw3540xkErESRQ5dsblgmCXQUKzdRs+9fMUFMum8s7AoMlFlth602WVjOUhKwT8IvWUjfplxBZrahZi///Xi8h1I7i20qVQwKGYM8JCf+5Un9bQxeQI1d8a4JZXqedVPJixIoS1VdF/JKiyGaaHm7Jq+U/sSw7aJP7tKW0W29E7bmG5nT01YEu+nz9XmznvO9x2OXn+ioUO1nbte7M5GLPproga1lB3QUPgtaFu1BiM7mxauGcW+nahp671joXhgqo7R51kWWk/sDca54yZFnR0NDLhrbV9uSQsyzqtAZn4mDFOpriYKvQ+p+se1rDy7irFp8NmQ0vasdm0gdX17n6r7mtQfmukZA6m5454o8a5IKDG6IBv6UbLFfH37WS0u0/MXiEOeZcvlMq1smzquso1fyLDy5uwyPU/r2Jhe4MRNuJlvO21f8BKrijjVLutNMuFaRyMm0z2BkIC01sDAeXohIb0LsUE7AnlNUY2WkDTIZg0dzNydtN6y/zYtEenfmHmD2vaDT7hZb3TzCFvdyBobkOS7vTYWzywBuUJxWa/lUh7YdJ18/tYSi4hnCSyQtdzlsEl1kOcS8jmaQD8o8ae7zQL9We0X7knw2ya3K+EbTStvkMBXWu0Xcjfrkq1cBMhweFUU5OPI7WhkdePpcv3xHhLAjd734pJgyfZBop+EdCjOAYL8dskrLuv1IN2u29kPR6967CbCYC0Mzbqu+w+2WC0H",sidebar_class_name:"get api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Add Permission to Role",permalink:"/docs/singulatron/add-permission-to-role"},next:{title:"Set Role Permissions",permalink:"/docs/singulatron/set-role-permission"}},N={},I=[];function w(e){const s={p:"p",...(0,n.R)(),...e.components},{Details:i}=s;return i||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(v.default,{as:"h1",className:"openapi__heading",children:"Get Permissions by Role"}),"\n",(0,a.jsx)(d(),{method:"get",path:"/user-svc/role/{roleId}/permissions"}),"\n",(0,a.jsx)(s.p,{children:"Retrieve permissions associated with a specific role ID."}),"\n",(0,a.jsx)(v.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,a.jsxs)(i,{style:{marginBottom:"1rem"},className:"openapi-markdown__details","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},children:(0,a.jsx)("h3",{className:"openapi-markdown__details-summary-header-params",children:(0,a.jsx)(s.p,{children:"Path Parameters"})})}),(0,a.jsx)("div",{children:(0,a.jsx)("ul",{children:(0,a.jsx)(p(),{className:"paramsItem",param:{description:"Role ID",in:"path",name:"roleId",required:!0,schema:{type:"integer"}}})})})]}),"\n",(0,a.jsx)("div",{children:(0,a.jsx)("div",{children:(0,a.jsxs)(l(),{label:void 0,id:void 0,children:[(0,a.jsxs)(f.default,{label:"200",value:"200",children:[(0,a.jsx)("div",{children:(0,a.jsx)(s.p,{children:"OK"})}),(0,a.jsx)("div",{children:(0,a.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,a.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,a.jsxs)(g(),{className:"openapi-tabs__schema",children:[(0,a.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,a.jsx)("strong",{children:(0,a.jsx)(s.p,{children:"Schema"})})}),(0,a.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,a.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,a.jsx)(j(),{collapsible:!0,className:"schemaItem",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details",children:[(0,a.jsx)("summary",{style:{},children:(0,a.jsxs)("span",{className:"openapi-schema__container",children:[(0,a.jsx)("strong",{className:"openapi-schema__property",children:(0,a.jsx)(s.p,{children:"permissions"})}),(0,a.jsx)("span",{className:"openapi-schema__name",children:(0,a.jsx)(s.p,{children:"object[]"})})]})}),(0,a.jsxs)("div",{style:{marginLeft:"1rem"},children:[(0,a.jsx)("li",{children:(0,a.jsx)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem",paddingBottom:".5rem"},children:(0,a.jsx)(s.p,{children:"Array ["})})}),(0,a.jsx)(j(),{collapsible:!1,name:"createdAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,a.jsx)(j(),{collapsible:!1,name:"description",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,a.jsx)(j(),{collapsible:!1,name:"id",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:'eg. "user.viewer"',type:"string"}}),(0,a.jsx)(j(),{collapsible:!1,name:"name",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:'eg. "User Viewer"',type:"string"}}),(0,a.jsx)(j(),{collapsible:!1,name:"ownerId",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:"Service who owns the permission",type:"string"}}),(0,a.jsx)(j(),{collapsible:!1,name:"updatedAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,a.jsx)("li",{children:(0,a.jsx)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem"},children:(0,a.jsx)(s.p,{children:"]"})})})]})]})})})]})}),(0,a.jsx)(f.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,a.jsx)(u(),{responseExample:'{\n  "permissions": [\n    {\n      "createdAt": "string",\n      "description": "string",\n      "id": "string",\n      "name": "string",\n      "ownerId": "string",\n      "updatedAt": "string"\n    }\n  ]\n}',language:"json"})})]})})})})]}),(0,a.jsxs)(f.default,{label:"400",value:"400",children:[(0,a.jsx)("div",{children:(0,a.jsx)(s.p,{children:"Invalid JSON"})}),(0,a.jsx)("div",{children:(0,a.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,a.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,a.jsx)(g(),{className:"openapi-tabs__schema",children:(0,a.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,a.jsx)("strong",{children:(0,a.jsx)(s.p,{children:"Schema"})})}),(0,a.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,a.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,a.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,a.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,a.jsxs)(f.default,{label:"401",value:"401",children:[(0,a.jsx)("div",{children:(0,a.jsx)(s.p,{children:"Unauthorized"})}),(0,a.jsx)("div",{children:(0,a.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,a.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,a.jsx)(g(),{className:"openapi-tabs__schema",children:(0,a.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,a.jsx)("strong",{children:(0,a.jsx)(s.p,{children:"Schema"})})}),(0,a.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,a.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,a.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,a.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,a.jsxs)(f.default,{label:"500",value:"500",children:[(0,a.jsx)("div",{children:(0,a.jsx)(s.p,{children:"Internal Server Error"})}),(0,a.jsx)("div",{children:(0,a.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,a.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,a.jsx)(g(),{className:"openapi-tabs__schema",children:(0,a.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,a.jsxs)(i,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,a.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,a.jsx)("strong",{children:(0,a.jsx)(s.p,{children:"Schema"})})}),(0,a.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,a.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,a.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,a.jsx)(s.p,{children:"string"})})})]})})})})})})]})]})})})]})}function k(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,a.jsx)(s,{...e,children:(0,a.jsx)(w,{...e})}):w(e)}}}]);