"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2183],{43934:(e,a,s)=>{s.r(a),s.d(a,{assets:()=>k,contentTitle:()=>y,default:()=>_,frontMatter:()=>w,metadata:()=>S,toc:()=>D});var l=s(74848),n=s(28453),t=s(91366),i=s.n(t),r=(s(6050),s(57742)),d=s.n(r),o=(s(67792),s(27362)),m=s.n(o),c=s(36683),h=s.n(c),p=s(81124),u=s.n(p),x=s(60674),j=s.n(x),f=s(23397),b=s.n(f),g=(s(26651),s(51107)),v=(s(77675),s(19365));const w={id:"make-default",title:"Make a Model Default",description:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.",sidebar_label:"Make a Model Default",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU1v3DYQ/SvEnFqAlhy7AQqd6iJtsE1SG1n7UDh7GEuzEmOKZEhqtxuBQH9Ef2F+STCS9sP2ukAvvuxK4szwzeN7wx4qCqVXLiproIA5xSBQtLYiLTCI2JCoaImdjtPHb//8K9YNGeG8bR1HexKBTBRrFRvbRYHiwxA5eyOP5KsgukBVBhKsI4+88ayCAlq8pzdjKEhw6LGlSD5Acds/QrmtDxIUvzuMDUgw2BIX4tVZBRI8femUpwqK6DuSEMqGWoSih7hxHBqiV6aGlBYcHJw1gQKvn52e8t/DbS/fgYTSmkgm8io6p1U5dJB/DhzSP93C3n2mMkJKKUn46VjZmVmhVpX4Y3755//ZwHkmMKoRMXlv/bHW5DNIXj1FcmOwi4316itVL4bk9XFOInmDWszJr8iL34aaLwMpSQhUdl7FzaC9Xwk9+YsuNlDcLlgqEWuW5aRDhqhKgoWElmJjWcyuG0XMOZAPijwJq3J8yvtJoiln0Z9UO9WHodtR8p3XnJtrW6JubIjF65/Pzl8BA9jim3O/Y4uHKB+zeb1xJD5NIZ9ALK3Wdk2VuNsIFMFhSQJNJaK9JyOwHF0jlt62g4NvAnkRpj7Fe1srI8hUzioTs60JG8KK/N6GF5OUhhOCHdHo1DvaDNQrs7QMlg8Vy+FQqUXFbQfUFH4JytSdxuityUrbHtS+mol555z1zNrIVBOjK/I8dI680xiX1rcZqhySfETHx84M7Va0Im2dCKSXJ8wwVeJiJtC5kIm/bOd5xtUe2xbvNAllThrbBRJvr64zcd2Q+F15usNAYmn9QBRn18SUaFWSCcQ9bTG/vXovVufZ6QPEocjz9Xqd1abLrK/zKS/kWDt9cp6dZk1sNfcQybfhcrlV277hNdY1+UzZfAjJmWsVNYfM9wSCBJbWyMBpdsYlnQ2xRXMA8gPe026A76fxA/76vQtf9r6YFBTp75g7jcpwDwOV/WS1W9hZDeT4DBKK/Y3wwG8LCXzqnNb3fI43XqfEn7905Nn8Cwkr9IqPf7yGVODnCool6kD/QcwPH6e750dxcFsd7WBrDLPhM0Ld8RtIuKfNwW2WFkluPcZQxtWLsiQXD/KeTMV0OJaubq5BAk5DYu9ILia3D1z9KKbHjh4h8G+Sz6T0/ej3lHbx49KzGbsxMkYzRYuU0ncxQ/e4",sidebar_class_name:"put api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},y=void 0,S={id:"singulatron/make-default",title:"Make a Model Default",description:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.",source:"@site/docs/singulatron/make-default.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/make-default",permalink:"/docs/singulatron/make-default",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"make-default",title:"Make a Model Default",description:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used.",sidebar_label:"Make a Model Default",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU1v3DYQ/SvEnFqAlhy7AQqd6iJtsE1SG1n7UDh7GEuzEmOKZEhqtxuBQH9Ef2F+STCS9sP2ukAvvuxK4szwzeN7wx4qCqVXLiproIA5xSBQtLYiLTCI2JCoaImdjtPHb//8K9YNGeG8bR1HexKBTBRrFRvbRYHiwxA5eyOP5KsgukBVBhKsI4+88ayCAlq8pzdjKEhw6LGlSD5Acds/QrmtDxIUvzuMDUgw2BIX4tVZBRI8femUpwqK6DuSEMqGWoSih7hxHBqiV6aGlBYcHJw1gQKvn52e8t/DbS/fgYTSmkgm8io6p1U5dJB/DhzSP93C3n2mMkJKKUn46VjZmVmhVpX4Y3755//ZwHkmMKoRMXlv/bHW5DNIXj1FcmOwi4316itVL4bk9XFOInmDWszJr8iL34aaLwMpSQhUdl7FzaC9Xwk9+YsuNlDcLlgqEWuW5aRDhqhKgoWElmJjWcyuG0XMOZAPijwJq3J8yvtJoiln0Z9UO9WHodtR8p3XnJtrW6JubIjF65/Pzl8BA9jim3O/Y4uHKB+zeb1xJD5NIZ9ALK3Wdk2VuNsIFMFhSQJNJaK9JyOwHF0jlt62g4NvAnkRpj7Fe1srI8hUzioTs60JG8KK/N6GF5OUhhOCHdHo1DvaDNQrs7QMlg8Vy+FQqUXFbQfUFH4JytSdxuityUrbHtS+mol555z1zNrIVBOjK/I8dI680xiX1rcZqhySfETHx84M7Va0Im2dCKSXJ8wwVeJiJtC5kIm/bOd5xtUe2xbvNAllThrbBRJvr64zcd2Q+F15usNAYmn9QBRn18SUaFWSCcQ9bTG/vXovVufZ6QPEocjz9Xqd1abLrK/zKS/kWDt9cp6dZk1sNfcQybfhcrlV277hNdY1+UzZfAjJmWsVNYfM9wSCBJbWyMBpdsYlnQ2xRXMA8gPe026A76fxA/76vQtf9r6YFBTp75g7jcpwDwOV/WS1W9hZDeT4DBKK/Y3wwG8LCXzqnNb3fI43XqfEn7905Nn8Cwkr9IqPf7yGVODnCool6kD/QcwPH6e750dxcFsd7WBrDLPhM0Ld8RtIuKfNwW2WFkluPcZQxtWLsiQXD/KeTMV0OJaubq5BAk5DYu9ILia3D1z9KKbHjh4h8G+Sz6T0/ej3lHbx49KzGbsxMkYzRYuU0ncxQ/e4",sidebar_class_name:"put api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Get a Model",permalink:"/docs/singulatron/get-model"},next:{title:"Start a Model",permalink:"/docs/singulatron/start-model"}},k={},D=[];function N(e){const a={p:"p",...(0,n.R)(),...e.components},{Details:s}=a;return s||function(e,a){throw new Error("Expected "+(a?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Make a Model Default"}),"\n",(0,l.jsx)(d(),{method:"put",path:"/model-svc/model/{modelId}/make-default"}),"\n",(0,l.jsx)(a.p,{children:"Sets a model as the default model \u2014 when prompts are sent without a Model ID, the default model is used."}),"\n",(0,l.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,l.jsxs)(s,{style:{marginBottom:"1rem"},className:"openapi-markdown__details","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},children:(0,l.jsx)("h3",{className:"openapi-markdown__details-summary-header-params",children:(0,l.jsx)(a.p,{children:"Path Parameters"})})}),(0,l.jsx)("div",{children:(0,l.jsx)("ul",{children:(0,l.jsx)(h(),{className:"paramsItem",param:{description:"Model ID",in:"path",name:"modelId",required:!0,schema:{type:"string"}}})})})]}),"\n",(0,l.jsx)("div",{children:(0,l.jsx)("div",{children:(0,l.jsxs)(i(),{label:void 0,id:void 0,children:[(0,l.jsxs)(v.default,{label:"200",value:"200",children:[(0,l.jsx)("div",{children:(0,l.jsx)(a.p,{children:"OK"})}),(0,l.jsx)("div",{children:(0,l.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,l.jsxs)(b(),{className:"openapi-tabs__schema",children:[(0,l.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(s,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(a.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,l.jsx)(a.p,{children:"object"})})})]})}),(0,l.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,l.jsx)(u(),{responseExample:"{}",language:"json"})})]})})})})]}),(0,l.jsxs)(v.default,{label:"400",value:"400",children:[(0,l.jsx)("div",{children:(0,l.jsx)(a.p,{children:"Invalid JSON"})}),(0,l.jsx)("div",{children:(0,l.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,l.jsxs)(b(),{className:"openapi-tabs__schema",children:[(0,l.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(s,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(a.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)(j(),{collapsible:!1,name:"error",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})}),(0,l.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,l.jsx)(u(),{responseExample:'{\n  "error": "string"\n}',language:"json"})})]})})})})]}),(0,l.jsxs)(v.default,{label:"401",value:"401",children:[(0,l.jsx)("div",{children:(0,l.jsx)(a.p,{children:"Unauthorized"})}),(0,l.jsx)("div",{children:(0,l.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,l.jsxs)(b(),{className:"openapi-tabs__schema",children:[(0,l.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(s,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(a.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)(j(),{collapsible:!1,name:"error",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})}),(0,l.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,l.jsx)(u(),{responseExample:'{\n  "error": "string"\n}',language:"json"})})]})})})})]}),(0,l.jsxs)(v.default,{label:"500",value:"500",children:[(0,l.jsx)("div",{children:(0,l.jsx)(a.p,{children:"Internal Server Error"})}),(0,l.jsx)("div",{children:(0,l.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(v.default,{label:"application/json",value:"application/json",children:(0,l.jsxs)(b(),{className:"openapi-tabs__schema",children:[(0,l.jsx)(v.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(s,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(a.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)(j(),{collapsible:!1,name:"error",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})}),(0,l.jsx)(v.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,l.jsx)(u(),{responseExample:'{\n  "error": "string"\n}',language:"json"})})]})})})})]})]})})})]})}function _(e={}){const{wrapper:a}={...(0,n.R)(),...e.components};return a?(0,l.jsx)(a,{...e,children:(0,l.jsx)(N,{...e})}):N(e)}}}]);