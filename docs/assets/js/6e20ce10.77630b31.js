"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6149],{27295:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>N,contentTitle:()=>D,default:()=>M,frontMatter:()=>v,metadata:()=>_,toc:()=>I});var l=a(74848),n=a(28453),i=a(91366),r=a.n(i),t=(a(6050),a(57742)),d=a.n(t),c=(a(67792),a(27362)),o=a.n(c),m=a(36683),p=a.n(m),h=a(81124),x=a.n(h),u=a(60674),j=a.n(u),g=a(23397),y=a.n(g),b=(a(26651),a(51107)),f=(a(77675),a(19365));const v={id:"delete-message",title:"Delete a Message",description:"Delete a specific message from a chat thread by its ID",sidebar_label:"Delete a Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVU1v3DYQ/SvEnFpAljZ2AxQ61a3dYFu3MbL2oXD2MKZmJSYUyZDUbrYC/3sxkvbDm02BtJf9EGf43rx5M+qhoiC9clFZAyXckKZIAkVwJNVKSdFSCFiTWHnbChSywShi4wkr8bwVKgYxv4EMrCOPfMm8ghKq4Zo/xlTIwKHHliL5AOVTf4I5hY33KH7iMDaQgcGWoISJwbyCDDx96pSnCsroO8ogyIZahLKHuHUcHKJXpoaUlhwcnDWBAp9fzmb8dR46dFJSCKtO660YyTOatCaSiZyHzmklhwqLD4GT+yNwrCrFR6jvPSsRFYOOFCdi9vkDyQgppZTBD+fIzM0atarEb4u3f34L+GnlI8CrLwEeDXaxsV79/W3VnQN4fb6CSN6gFgvya/Li1nvr/x9SyiCQ7LyK28E6PxN68tddbKB8WnKXI9bsKviFncnAShIsM2gpNvZgxsGFnAUFe/girGUxOavo9xZLwHhMfnRq5zVnFNpK1I0NsXz94+XVK2DgHa8F0x9NdszuVJyHrSPxfgp5D2JltbYbGsaI5w0lCTSViPYjGYFyNPo4d7Eh8RjIizAVKO5srYwgUzmrTMx3k9MQVuQPs3M9NXwQHPZmRKd+py2wusqsLJPlHqEcekQtKi47oKbwU1Cm7jRGb00ubXt09/1cLDrnrI+QTUo1MbqyKELnyDuNcWV9m6MqIGUncrzrzFBuRWvS1olAenXBClMlrucCnQu5+Mt2Xjhva49ti8+ahDIXje0CiTf3D7l4aEj8qjw9YyCxsn4QirNrYkm0kmQCcU07zm/u78T6Kp+9YBzKothsNnltutz6upjyQoG10xdX+SxvYqu5hki+DW9XO5sdCt5gXZPPlS2GkIK1VlFzyOIgIGTA1hoVmOWXfKWzIbZojkjul/Bhgb7Qrj8M1H9f2JMTIn2OhdOoDHMZJOmnQXmC3aBAtlvCkEF52MfLDLhjHNr33INHr1Pix5868jywywzW6BW3btz8KvDvCsoV6kD/Uth376ZV/7148YI4y3tna7NlhVF3/A8y+EjbFy+QtEzZbkaYznh+LSW5eJT5xZJKx/vk5vbu9uEWMsBpzg9Dxfdlux8McJbY6VCOLPgzZV9J6ftxZFPax49HX83Yb4IxmnVappT+AWa0tzw=",sidebar_class_name:"delete api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},D=void 0,_={id:"singulatron/delete-message",title:"Delete a Message",description:"Delete a specific message from a chat thread by its ID",source:"@site/docs/singulatron/delete-message.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/delete-message",permalink:"/docs/singulatron/delete-message",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"delete-message",title:"Delete a Message",description:"Delete a specific message from a chat thread by its ID",sidebar_label:"Delete a Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVU1v3DYQ/SvEnFpAljZ2AxQ61a3dYFu3MbL2oXD2MKZmJSYUyZDUbrYC/3sxkvbDm02BtJf9EGf43rx5M+qhoiC9clFZAyXckKZIAkVwJNVKSdFSCFiTWHnbChSywShi4wkr8bwVKgYxv4EMrCOPfMm8ghKq4Zo/xlTIwKHHliL5AOVTf4I5hY33KH7iMDaQgcGWoISJwbyCDDx96pSnCsroO8ogyIZahLKHuHUcHKJXpoaUlhwcnDWBAp9fzmb8dR46dFJSCKtO660YyTOatCaSiZyHzmklhwqLD4GT+yNwrCrFR6jvPSsRFYOOFCdi9vkDyQgppZTBD+fIzM0atarEb4u3f34L+GnlI8CrLwEeDXaxsV79/W3VnQN4fb6CSN6gFgvya/Li1nvr/x9SyiCQ7LyK28E6PxN68tddbKB8WnKXI9bsKviFncnAShIsM2gpNvZgxsGFnAUFe/girGUxOavo9xZLwHhMfnRq5zVnFNpK1I0NsXz94+XVK2DgHa8F0x9NdszuVJyHrSPxfgp5D2JltbYbGsaI5w0lCTSViPYjGYFyNPo4d7Eh8RjIizAVKO5srYwgUzmrTMx3k9MQVuQPs3M9NXwQHPZmRKd+py2wusqsLJPlHqEcekQtKi47oKbwU1Cm7jRGb00ubXt09/1cLDrnrI+QTUo1MbqyKELnyDuNcWV9m6MqIGUncrzrzFBuRWvS1olAenXBClMlrucCnQu5+Mt2Xjhva49ti8+ahDIXje0CiTf3D7l4aEj8qjw9YyCxsn4QirNrYkm0kmQCcU07zm/u78T6Kp+9YBzKothsNnltutz6upjyQoG10xdX+SxvYqu5hki+DW9XO5sdCt5gXZPPlS2GkIK1VlFzyOIgIGTA1hoVmOWXfKWzIbZojkjul/Bhgb7Qrj8M1H9f2JMTIn2OhdOoDHMZJOmnQXmC3aBAtlvCkEF52MfLDLhjHNr33INHr1Pix5868jywywzW6BW3btz8KvDvCsoV6kD/Uth376ZV/7148YI4y3tna7NlhVF3/A8y+EjbFy+QtEzZbkaYznh+LSW5eJT5xZJKx/vk5vbu9uEWMsBpzg9Dxfdlux8McJbY6VCOLPgzZV9J6ftxZFPax49HX83Yb4IxmnVappT+AWa0tzw=",sidebar_class_name:"delete api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Introduction",permalink:"/docs/singulatron/singulatron"},next:{title:"Add Thread",permalink:"/docs/singulatron/add-thread"}},N={},I=[];function w(e){const s={p:"p",...(0,n.R)(),...e.components},{Details:a}=s;return a||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Delete a Message"}),"\n",(0,l.jsx)(d(),{method:"delete",path:"/chat-svc/message/{messageId}"}),"\n",(0,l.jsx)(s.p,{children:"Delete a specific message from a chat thread by its ID"}),"\n",(0,l.jsx)(b.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,l.jsxs)(a,{style:{marginBottom:"1rem"},className:"openapi-markdown__details","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},children:(0,l.jsx)("h3",{className:"openapi-markdown__details-summary-header-params",children:(0,l.jsx)(s.p,{children:"Path Parameters"})})}),(0,l.jsx)("div",{children:(0,l.jsx)("ul",{children:(0,l.jsx)(p(),{className:"paramsItem",param:{description:"Message ID",in:"path",name:"messageId",required:!0,schema:{type:"string"}}})})})]}),"\n",(0,l.jsx)("div",{children:(0,l.jsx)("div",{children:(0,l.jsxs)(r(),{label:void 0,id:void 0,children:[(0,l.jsxs)(f.default,{label:"200",value:"200",children:[(0,l.jsx)("div",{children:(0,l.jsx)(s.p,{children:"Message successfully deleted"})}),(0,l.jsx)("div",{children:(0,l.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,l.jsxs)(y(),{className:"openapi-tabs__schema",children:[(0,l.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(s.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)(j(),{name:"property name*",required:!1,schemaName:"any",qualifierMessage:void 0,schema:{additionalProperties:!0,type:"object"},collapsible:!1,discriminator:!1})})]})}),(0,l.jsx)(f.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,l.jsx)(x(),{responseExample:"{}",language:"json"})})]})})})})]}),(0,l.jsxs)(f.default,{label:"400",value:"400",children:[(0,l.jsx)("div",{children:(0,l.jsx)(s.p,{children:"Invalid JSON"})}),(0,l.jsx)("div",{children:(0,l.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,l.jsx)(y(),{className:"openapi-tabs__schema",children:(0,l.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(s.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,l.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,l.jsxs)(f.default,{label:"401",value:"401",children:[(0,l.jsx)("div",{children:(0,l.jsx)(s.p,{children:"Unauthorized"})}),(0,l.jsx)("div",{children:(0,l.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,l.jsx)(y(),{className:"openapi-tabs__schema",children:(0,l.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(s.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,l.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,l.jsxs)(f.default,{label:"500",value:"500",children:[(0,l.jsx)("div",{children:(0,l.jsx)(s.p,{children:"Internal Server Error"})}),(0,l.jsx)("div",{children:(0,l.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,l.jsx)(y(),{className:"openapi-tabs__schema",children:(0,l.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(s.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,l.jsx)(s.p,{children:"string"})})})]})})})})})})]})]})})})]})}function M(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,l.jsx)(s,{...e,children:(0,l.jsx)(w,{...e})}):w(e)}}}]);