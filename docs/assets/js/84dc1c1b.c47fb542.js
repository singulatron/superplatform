"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5933],{28455:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>w,contentTitle:()=>b,default:()=>S,frontMatter:()=>f,metadata:()=>y,toc:()=>_});var r=a(74848),i=a(28453),l=a(91366),n=a.n(l),t=(a(6050),a(57742)),d=a.n(t),o=(a(67792),a(27362)),c=a.n(o),m=(a(36683),a(81124)),p=a.n(m),h=a(60674),x=a.n(h),j=a(23397),u=a.n(j),g=(a(26651),a(51107)),v=(a(77675),a(19365));const f={id:"register",title:"Register a New User",description:"Register a new user with a name, email, and password.",sidebar_label:"Register a New User",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VVFv2zgM/isCnx07127A4KfrDtvQbViLpns4dH1gbMbWJksaJcfLBf7vA+W4TbvesN3D5SUSJVIfP/Kj91BTqFj7qJ2FEq6o0SESK1SWBtUHYjXo2MoeO8oUdahNptDWymMIg+M6hwycJ0aJcV5DCXyIAhkwfe0pxJeu3kG5h8rZSDbKEr03ukpOxecgr+8hVC11KCvPEjJqCrJLr8oi7jxBCSGytg2MGQiqJw9mdE8cjtlscevPVEUYxfQvRFxNCRxS0Uw1lJF7GsUQvLNhgniyXMrfwygX7yD79Zx/ADVm8OypsOd2i0bX6u3q4sPvPPCIVGbHv0jPmMHzp5FEYotGrYi3xOpVivn/QBIjNgHKG/gofSoQdEVwm0FHsXXSid6l0nmMLZRQSD8vwrYqjjo0JOQSZg89G7lWGFehaV2I5fMXJ6d/wHgr96qeddytBPsE9yUhE5/1EvwxM9c7T+rT4conUBtnjBuoVuudQhU8VpRUFN0XsgqrqbfUhl2nYksqpRQOOan3rtFWka290zaK4rQ80hLWKYlJBiBQHOt/EttwRxp6/Y52iUapwNW9Jl99w84bOtLYzPldzLv9vaKOZKbtxs3Cxio+CISGwp9B26Y3GNnZvHLdEdbLc7XqvXcsFZqYb2P0ZVGE3hN7g3HjuMtRF/CjPnub6KtpS8Z5FchsFlIxqtXZuULvQ67+dj0rz65h7DpcG1LaLlrXB1JvLq9zdd2Seq2Z1hhIbRwn4sW7IaHY6IpsSOTMmN9cvlfb03z5AHEoi2IYhryxfe64KQ5+ocDGm8Vpvszb2BnJIRJ34WIzd+p9wgM2DXGuXZGuFFI7HaUwsLonEDKQVp0YWOYnac65EDu0RyCPZvgHGlIjwSP6jgbxfx75h96K9C0W3qC2giaRsj/o7QZmvaXpeVDcbQZSJzne74X5j2zGUcxfe+IdlDe3GWyRtRRMdmM297lI9AvtoIS/JvgLUZmwgqafGv3RrBmz2eOsqsjHn949HhyXF6tryGB9+G51rhYfxkFSwQFKSF898U6jINn2YNA2PTZyd4opv+9vhIvk",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},b=void 0,y={id:"singulatron/register",title:"Register a New User",description:"Register a new user with a name, email, and password.",source:"@site/docs/singulatron/register.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/register",permalink:"/docs/singulatron/register",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"register",title:"Register a New User",description:"Register a new user with a name, email, and password.",sidebar_label:"Register a New User",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VVFv2zgM/isCnx07127A4KfrDtvQbViLpns4dH1gbMbWJksaJcfLBf7vA+W4TbvesN3D5SUSJVIfP/Kj91BTqFj7qJ2FEq6o0SESK1SWBtUHYjXo2MoeO8oUdahNptDWymMIg+M6hwycJ0aJcV5DCXyIAhkwfe0pxJeu3kG5h8rZSDbKEr03ukpOxecgr+8hVC11KCvPEjJqCrJLr8oi7jxBCSGytg2MGQiqJw9mdE8cjtlscevPVEUYxfQvRFxNCRxS0Uw1lJF7GsUQvLNhgniyXMrfwygX7yD79Zx/ADVm8OypsOd2i0bX6u3q4sPvPPCIVGbHv0jPmMHzp5FEYotGrYi3xOpVivn/QBIjNgHKG/gofSoQdEVwm0FHsXXSid6l0nmMLZRQSD8vwrYqjjo0JOQSZg89G7lWGFehaV2I5fMXJ6d/wHgr96qeddytBPsE9yUhE5/1EvwxM9c7T+rT4conUBtnjBuoVuudQhU8VpRUFN0XsgqrqbfUhl2nYksqpRQOOan3rtFWka290zaK4rQ80hLWKYlJBiBQHOt/EttwRxp6/Y52iUapwNW9Jl99w84bOtLYzPldzLv9vaKOZKbtxs3Cxio+CISGwp9B26Y3GNnZvHLdEdbLc7XqvXcsFZqYb2P0ZVGE3hN7g3HjuMtRF/CjPnub6KtpS8Z5FchsFlIxqtXZuULvQ67+dj0rz65h7DpcG1LaLlrXB1JvLq9zdd2Seq2Z1hhIbRwn4sW7IaHY6IpsSOTMmN9cvlfb03z5AHEoi2IYhryxfe64KQ5+ocDGm8Vpvszb2BnJIRJ34WIzd+p9wgM2DXGuXZGuFFI7HaUwsLonEDKQVp0YWOYnac65EDu0RyCPZvgHGlIjwSP6jgbxfx75h96K9C0W3qC2giaRsj/o7QZmvaXpeVDcbQZSJzne74X5j2zGUcxfe+IdlDe3GWyRtRRMdmM297lI9AvtoIS/JvgLUZmwgqafGv3RrBmz2eOsqsjHn949HhyXF6tryGB9+G51rhYfxkFSwQFKSF898U6jINn2YNA2PTZyd4opv+9vhIvk",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Ge Public Key",permalink:"/docs/singulatron/get-public-key"},next:{title:"Create a New Role",permalink:"/docs/singulatron/create-role"}},w={},_=[];function N(e){const s={p:"p",...(0,i.R)(),...e.components},{Details:a}=s;return a||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Register a New User"}),"\n",(0,r.jsx)(d(),{method:"post",path:"/user-svc/register"}),"\n",(0,r.jsx)(s.p,{children:"Register a new user with a name, email, and password."}),"\n",(0,r.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,r.jsx)(c(),{className:"openapi-tabs__mime",children:(0,r.jsx)(v.default,{label:"application/json",value:"application/json-schema",children:(0,r.jsxs)(a,{style:{},className:"openapi-markdown__details mime","data-collapsed":!1,open:!0,children:[(0,r.jsxs)("summary",{style:{},className:"openapi-markdown__details-summary-mime",children:[(0,r.jsx)("h3",{className:"openapi-markdown__details-summary-header-body",children:(0,r.jsx)(s.p,{children:"Body"})}),(0,r.jsx)("strong",{className:"openapi-schema__required",children:(0,r.jsx)(s.p,{children:"required"})})]}),(0,r.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"},children:(0,r.jsx)("div",{style:{marginTop:"1rem",marginBottom:"1rem"},children:(0,r.jsx)(s.p,{children:"Register Request"})})}),(0,r.jsxs)("ul",{style:{marginLeft:"1rem"},children:[(0,r.jsx)(x(),{collapsible:!1,name:"email",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,r.jsx)(x(),{collapsible:!1,name:"name",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,r.jsx)(x(),{collapsible:!1,name:"password",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})]})]})})}),"\n",(0,r.jsx)("div",{children:(0,r.jsx)("div",{children:(0,r.jsxs)(n(),{label:void 0,id:void 0,children:[(0,r.jsxs)(v.default,{label:"200",value:"200",children:[(0,r.jsx)("div",{children:(0,r.jsx)(s.p,{children:"OK"})}),(0,r.jsx)("div",{children:(0,r.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,r.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,r.jsxs)(u(),{className:"openapi-tabs__schema",children:[(0,r.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,r.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,r.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,r.jsx)("strong",{children:(0,r.jsx)(s.p,{children:"Schema"})})}),(0,r.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,r.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,r.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,r.jsx)(s.p,{children:"object"})})})]})}),(0,r.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,r.jsx)(p(),{responseExample:"{}",language:"json"})})]})})})})]}),(0,r.jsxs)(v.default,{label:"400",value:"400",children:[(0,r.jsx)("div",{children:(0,r.jsx)(s.p,{children:"Invalid JSON"})}),(0,r.jsx)("div",{children:(0,r.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,r.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,r.jsxs)(u(),{className:"openapi-tabs__schema",children:[(0,r.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,r.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,r.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,r.jsx)("strong",{children:(0,r.jsx)(s.p,{children:"Schema"})})}),(0,r.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,r.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,r.jsx)(x(),{collapsible:!1,name:"error",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})}),(0,r.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,r.jsx)(p(),{responseExample:'{\n  "error": "string"\n}',language:"json"})})]})})})})]}),(0,r.jsxs)(v.default,{label:"500",value:"500",children:[(0,r.jsx)("div",{children:(0,r.jsx)(s.p,{children:"Internal Server Error"})}),(0,r.jsx)("div",{children:(0,r.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,r.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,r.jsxs)(u(),{className:"openapi-tabs__schema",children:[(0,r.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,r.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,r.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,r.jsx)("strong",{children:(0,r.jsx)(s.p,{children:"Schema"})})}),(0,r.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,r.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,r.jsx)(x(),{collapsible:!1,name:"error",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})}),(0,r.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,r.jsx)(p(),{responseExample:'{\n  "error": "string"\n}',language:"json"})})]})})})})]})]})})})]})}function S(e={}){const{wrapper:s}={...(0,i.R)(),...e.components};return s?(0,r.jsx)(s,{...e,children:(0,r.jsx)(N,{...e})}):N(e)}}}]);