"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2731],{8742:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>_,contentTitle:()=>y,default:()=>q,frontMatter:()=>v,metadata:()=>b,toc:()=>N});var i=a(74848),n=a(28453),l=a(91366),r=a.n(l),t=(a(6050),a(57742)),d=a.n(t),c=(a(67792),a(27362)),o=a.n(c),p=(a(36683),a(81124)),h=a.n(p),m=a(60674),x=a.n(m),j=a(23397),u=a.n(j),g=(a(26651),a(51107)),f=(a(77675),a(19365));const v={id:"get-threads",title:"Get Threads",description:"Fetch all chat threads associated with a specific user",sidebar_label:"Get Threads",hide_title:!0,hide_table_of_contents:!0,api:"eJylVd1v2zYQ/1eIe1YkN1mBQU9LhzbIMCxB7D4MXR7O1FliR5HskbLrGf7fi6PkxHWMrt38Yup4H7/73Qd30FDUbEIy3kEN7yjpTqG1SneYVOqYsIkKY/TaYKJGbUzqFKoYSJuV0WqIxFAA06eBYnrjmy3UO9DeJXJJjhiCNRolQPUxSpQdRN1Rj3JK20BQg19+JJ1gv98XJ4huKKnFBONhDDKFM0wN1IkH2osgBu8iRXF6OZvJ39eODk7ioDXFuBqs3SqmxIbW1EDx/ZgD+0CczBhs4kiOJlEfX2poJqHuOh0lHBMb18K+ANOcFSeTLJ3JQsTKr1TqaCpPCcUZcx+Mvh1hnXiYblRDK+Moqk1ndKeyQTxyq5ZkvWujSr78yy3Ga41OsR/azm7VUjTlnAQPRpWwjWrl+dA2AuyJkpcIRwEy41a+h9B8gyVps7PpvB8vMnDTxAM1oi+p+Qw5kmA18Yix7wX2LDj06L9r5Db+6VwT3ro1WtOo3+Z3f/xIy52AnAK8OkOHwyF1ns0/P9bT5wK8Pp9BInZo1Zx4TazeMnv+f5GEQGwj1B9Atg48FtBT6nwDNQSf5z1g6qCGSu4vIvHaaKoOk1dAzFjEww4GtqJZWa/Rdj6m+vXPl1evYP8oenpgk7ZzQTMO5xuMRl8P4v4JW5dSEK9ZC2pYik4utKTz8Lzq3n7GPuQxlUF2K3/YfagzE9SjETQRLcVfonHtYDGxd6X2PRTgMPu/vr9V8yEEz5LrmIBgqKsqDoE4WEwrz32JpoIXG/JhcApdoxpak/VBRbKrC0mcGnV9qzCEWKo//cAqsG8Z+x6XlpRxF50fIqmb+0WpFh2pd4ZpiZGmGaZs3ZJMizWaXMz76ID55v53tb4qZ18hjnVVbTabsnVD6bmtJrtYYRvsxVU5K7vU2zxDxH28W83HWh4lvMG2JS6Nr7JKBU+7EObPBEIBUvGRgVl5KS6lV3p0RyCPHg84oe3ojfrvr97UL4k+pypYNC6vKmFjN7Xs2NKHlhWTCc5jAVIj0djthPX3bPd7EX8aiLdQf3gsYI1spFjytS+gI2yIc5f/TVuo4dcxhYuF4BB1O+TFdDqA++Jgca01hfRN3ePxu7+bL6CA5fSs974RG8aNvMG4gRqgAJ8pzdOUZTuw6NoBW9EdfcrvCwYh6q8=",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},y=void 0,b={id:"singulatron/get-threads",title:"Get Threads",description:"Fetch all chat threads associated with a specific user",source:"@site/docs/singulatron/get-threads.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/get-threads",permalink:"/docs/singulatron/get-threads",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"get-threads",title:"Get Threads",description:"Fetch all chat threads associated with a specific user",sidebar_label:"Get Threads",hide_title:!0,hide_table_of_contents:!0,api:"eJylVd1v2zYQ/1eIe1YkN1mBQU9LhzbIMCxB7D4MXR7O1FliR5HskbLrGf7fi6PkxHWMrt38Yup4H7/73Qd30FDUbEIy3kEN7yjpTqG1SneYVOqYsIkKY/TaYKJGbUzqFKoYSJuV0WqIxFAA06eBYnrjmy3UO9DeJXJJjhiCNRolQPUxSpQdRN1Rj3JK20BQg19+JJ1gv98XJ4huKKnFBONhDDKFM0wN1IkH2osgBu8iRXF6OZvJ39eODk7ioDXFuBqs3SqmxIbW1EDx/ZgD+0CczBhs4kiOJlEfX2poJqHuOh0lHBMb18K+ANOcFSeTLJ3JQsTKr1TqaCpPCcUZcx+Mvh1hnXiYblRDK+Moqk1ndKeyQTxyq5ZkvWujSr78yy3Ga41OsR/azm7VUjTlnAQPRpWwjWrl+dA2AuyJkpcIRwEy41a+h9B8gyVps7PpvB8vMnDTxAM1oi+p+Qw5kmA18Yix7wX2LDj06L9r5Db+6VwT3ro1WtOo3+Z3f/xIy52AnAK8OkOHwyF1ns0/P9bT5wK8Pp9BInZo1Zx4TazeMnv+f5GEQGwj1B9Atg48FtBT6nwDNQSf5z1g6qCGSu4vIvHaaKoOk1dAzFjEww4GtqJZWa/Rdj6m+vXPl1evYP8oenpgk7ZzQTMO5xuMRl8P4v4JW5dSEK9ZC2pYik4utKTz8Lzq3n7GPuQxlUF2K3/YfagzE9SjETQRLcVfonHtYDGxd6X2PRTgMPu/vr9V8yEEz5LrmIBgqKsqDoE4WEwrz32JpoIXG/JhcApdoxpak/VBRbKrC0mcGnV9qzCEWKo//cAqsG8Z+x6XlpRxF50fIqmb+0WpFh2pd4ZpiZGmGaZs3ZJMizWaXMz76ID55v53tb4qZ18hjnVVbTabsnVD6bmtJrtYYRvsxVU5K7vU2zxDxH28W83HWh4lvMG2JS6Nr7JKBU+7EObPBEIBUvGRgVl5KS6lV3p0RyCPHg84oe3ojfrvr97UL4k+pypYNC6vKmFjN7Xs2NKHlhWTCc5jAVIj0djthPX3bPd7EX8aiLdQf3gsYI1spFjytS+gI2yIc5f/TVuo4dcxhYuF4BB1O+TFdDqA++Jgca01hfRN3ePxu7+bL6CA5fSs974RG8aNvMG4gRqgAJ8pzdOUZTuw6NoBW9EdfcrvCwYh6q8=",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Get Messages",permalink:"/docs/singulatron/get-messages"},next:{title:"Get",permalink:"/docs/singulatron/get"}},_={},N=[];function T(e){const s={p:"p",...(0,n.R)(),...e.components},{Details:a}=s;return a||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Get Threads"}),"\n",(0,i.jsx)(d(),{method:"post",path:"/chat-service/threads"}),"\n",(0,i.jsx)(s.p,{children:"Fetch all chat threads associated with a specific user"}),"\n",(0,i.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(o(),{className:"openapi-tabs__mime",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json-schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details mime","data-collapsed":!1,open:!0,children:[(0,i.jsxs)("summary",{style:{},className:"openapi-markdown__details-summary-mime",children:[(0,i.jsx)("h3",{className:"openapi-markdown__details-summary-header-body",children:(0,i.jsx)(s.p,{children:"Body"})}),(0,i.jsx)("strong",{className:"openapi-schema__required",children:(0,i.jsx)(s.p,{children:"required"})})]}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:"1rem",marginBottom:"1rem"},children:(0,i.jsx)(s.p,{children:"Get Threads Request"})})}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"object"})})})]})})}),"\n",(0,i.jsx)("div",{children:(0,i.jsx)("div",{children:(0,i.jsxs)(r(),{label:void 0,id:void 0,children:[(0,i.jsxs)(f.default,{label:"200",value:"200",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Threads successfully retrieved"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsxs)(u(),{className:"openapi-tabs__schema",children:[(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(x(),{collapsible:!0,className:"schemaItem",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details",children:[(0,i.jsx)("summary",{style:{},children:(0,i.jsxs)("span",{className:"openapi-schema__container",children:[(0,i.jsx)("strong",{className:"openapi-schema__property",children:(0,i.jsx)(s.p,{children:"threads"})}),(0,i.jsx)("span",{className:"openapi-schema__name",children:(0,i.jsx)(s.p,{children:"object[]"})})]})}),(0,i.jsxs)("div",{style:{marginLeft:"1rem"},children:[(0,i.jsx)("li",{children:(0,i.jsx)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem",paddingBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"Array ["})})}),(0,i.jsx)(x(),{collapsible:!1,name:"createdAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"id",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"title",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:"Title of the thread.",type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"topicIds",required:!1,schemaName:"string[]",qualifierMessage:void 0,schema:{description:"TopicIds defines which topics the thread belongs to.\nTopics can roughly be thought of as tags for threads.",items:{type:"string"},type:"array"}}),(0,i.jsx)(x(),{collapsible:!1,name:"updatedAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"userIds",required:!1,schemaName:"string[]",qualifierMessage:void 0,schema:{description:"UserIds the ids of the users who can see this thread.",items:{type:"string"},type:"array"}}),(0,i.jsx)("li",{children:(0,i.jsx)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem"},children:(0,i.jsx)(s.p,{children:"]"})})})]})]})})})]})}),(0,i.jsx)(f.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,i.jsx)(h(),{responseExample:'{\n  "threads": [\n    {\n      "createdAt": "string",\n      "id": "string",\n      "title": "string",\n      "topicIds": [\n        "string"\n      ],\n      "updatedAt": "string",\n      "userIds": [\n        "string"\n      ]\n    }\n  ]\n}',language:"json"})})]})})})})]}),(0,i.jsxs)(f.default,{label:"400",value:"400",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Invalid JSON"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(u(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,i.jsxs)(f.default,{label:"401",value:"401",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Unauthorized"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(u(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,i.jsxs)(f.default,{label:"500",value:"500",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Internal Server Error"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(u(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]})]})})})]})}function q(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,i.jsx)(s,{...e,children:(0,i.jsx)(T,{...e})}):T(e)}}}]);