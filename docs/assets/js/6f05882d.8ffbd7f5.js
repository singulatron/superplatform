"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5480],{2397:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>N,contentTitle:()=>b,default:()=>q,frontMatter:()=>y,metadata:()=>v,toc:()=>_});var i=a(74848),n=a(28453),l=a(91366),r=a.n(l),t=(a(6050),a(57742)),d=a.n(t),c=(a(67792),a(27362)),o=a.n(c),m=(a(36683),a(81124)),p=a.n(m),h=a(60674),x=a.n(h),u=a(23397),j=a.n(u),g=(a(26651),a(51107)),f=(a(77675),a(19365));const y={id:"add-message",title:"Add Message",description:"Add a new message to a specific chat thread",sidebar_label:"Add Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVktvGzcQ/ivEXHpZ7zp2AxR7aZUiSVW0tWE5h8L2YbQc7TLlkgwfVhRB/70Y7sqS5UVQozotZ4bz/OajtiApNF65qKyBGmZSChSG1qKnELAlEa1AERw1aqUa0XQYRew8oYQCPH1JFOI7KzdQb6GxJpKJ/InOadUge60+B3a9hdB01CN/OW8d+ago8GmM9FKBIVCcy/x9kuaoEZJWylAQsSOBMWLT9WTicN5X0GEooQAVqc+u4sYR1BCiV6aFXbEXoPe44fNRHc/D/joohF09C0BtKe7hN1WIdYfxhyCS+/keoHgZqfGEkeQsTuah5HR6ud1z+TKf21FzklA5FTo5+Z3QKZCfivApy4UaOqqeIrG9WHdWrL2N9Cz4vflgvZjN9wK+qoJYKdLZEfUubiZSPAzCLj9TEyclu2ICsH+Oc7gZ4DgCU3mSUEefaMeC4KwJA64uzs9fVrp3ElLTUAirpPVGoJTEQP/PyEYpFatQXx9BmZOYqGVXwI9TqczNI2olxe+Lq79eE/y0oUOANxNTNZhiZ7369rrqpgK8na4gkjeoxYL8I3nx3nvr/18kBgO2Aeo7YBKChwJ6ip2VUIOzeeoOYwc1VKyvRvRVKLnEkPPg21tIXrNVpW2DurMh1m9/urh8A7sHtmuSV3Gz4EwGtLzDoJpZYtdPeXUxOvaaraCGJdtkwHIpNwdafP8Ve6fphOYOxHa3L/HhqDl72TO+OAiZJg6nAzscZEe7fiQcV/x44ZRZ2T1zY5PnQj0q7k9ATeGXoEybNEZvTdnYHgowmCueXc/FIjlnPXd+aCl3pa6qkBx5pzGurO9LVBW82NqbZAQaKSQ9krZOBNKrMx4FSWYOdC6U4m+bvHDeth77HpeahDJnnU2BxMfr21LcdiQ+KE9LDCRW1mcW4tsDA2rVkAm53/ucP17/IR4vy/NnGYe6qtbrddmaVFrfVuO9UGHr9NlleV52sdeZiMn34WrFmFYNHRW8xrYlXypbZZOKx6Iijx0WhwZCAYzBoQPn5QW7ZOT2aI6TPBAanLTt6IV95UM9wjbS11g5jcpkzucWbMetGbeqeIJpwVzGsOSxsH675UZ/8nq3Y/GXRH4D9d1DAY/oFc+HT7sCOkJJPq/aP7Q5PJtnt5wFm+uU39tTBtgV+xuzpiEXv2t7vP/XV4tbKGA5/g/preQ7Htf8FOAaan6Mbe5iXuks24JG06a8kjD45N+/GpIe8Q==",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},b=void 0,v={id:"singulatron/add-message",title:"Add Message",description:"Add a new message to a specific chat thread",source:"@site/docs/singulatron/add-message.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/add-message",permalink:"/docs/singulatron/add-message",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"add-message",title:"Add Message",description:"Add a new message to a specific chat thread",sidebar_label:"Add Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVktvGzcQ/ivEXHpZ7zp2AxR7aZUiSVW0tWE5h8L2YbQc7TLlkgwfVhRB/70Y7sqS5UVQozotZ4bz/OajtiApNF65qKyBGmZSChSG1qKnELAlEa1AERw1aqUa0XQYRew8oYQCPH1JFOI7KzdQb6GxJpKJ/InOadUge60+B3a9hdB01CN/OW8d+ago8GmM9FKBIVCcy/x9kuaoEZJWylAQsSOBMWLT9WTicN5X0GEooQAVqc+u4sYR1BCiV6aFXbEXoPe44fNRHc/D/joohF09C0BtKe7hN1WIdYfxhyCS+/keoHgZqfGEkeQsTuah5HR6ud1z+TKf21FzklA5FTo5+Z3QKZCfivApy4UaOqqeIrG9WHdWrL2N9Cz4vflgvZjN9wK+qoJYKdLZEfUubiZSPAzCLj9TEyclu2ICsH+Oc7gZ4DgCU3mSUEefaMeC4KwJA64uzs9fVrp3ElLTUAirpPVGoJTEQP/PyEYpFatQXx9BmZOYqGVXwI9TqczNI2olxe+Lq79eE/y0oUOANxNTNZhiZ7369rrqpgK8na4gkjeoxYL8I3nx3nvr/18kBgO2Aeo7YBKChwJ6ip2VUIOzeeoOYwc1VKyvRvRVKLnEkPPg21tIXrNVpW2DurMh1m9/urh8A7sHtmuSV3Gz4EwGtLzDoJpZYtdPeXUxOvaaraCGJdtkwHIpNwdafP8Ve6fphOYOxHa3L/HhqDl72TO+OAiZJg6nAzscZEe7fiQcV/x44ZRZ2T1zY5PnQj0q7k9ATeGXoEybNEZvTdnYHgowmCueXc/FIjlnPXd+aCl3pa6qkBx5pzGurO9LVBW82NqbZAQaKSQ9krZOBNKrMx4FSWYOdC6U4m+bvHDeth77HpeahDJnnU2BxMfr21LcdiQ+KE9LDCRW1mcW4tsDA2rVkAm53/ucP17/IR4vy/NnGYe6qtbrddmaVFrfVuO9UGHr9NlleV52sdeZiMn34WrFmFYNHRW8xrYlXypbZZOKx6Iijx0WhwZCAYzBoQPn5QW7ZOT2aI6TPBAanLTt6IV95UM9wjbS11g5jcpkzucWbMetGbeqeIJpwVzGsOSxsH675UZ/8nq3Y/GXRH4D9d1DAY/oFc+HT7sCOkJJPq/aP7Q5PJtnt5wFm+uU39tTBtgV+xuzpiEXv2t7vP/XV4tbKGA5/g/preQ7Htf8FOAaan6Mbe5iXuks24JG06a8kjD45N+/GpIe8Q==",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Introduction",permalink:"/docs/singulatron/singulatron"},next:{title:"Delete Message",permalink:"/docs/singulatron/delete-message"}},N={},_=[];function A(e){const s={p:"p",...(0,n.R)(),...e.components},{Details:a}=s;return a||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Add Message"}),"\n",(0,i.jsx)(d(),{method:"post",path:"/chat/message/add"}),"\n",(0,i.jsx)(s.p,{children:"Add a new message to a specific chat thread"}),"\n",(0,i.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(o(),{className:"openapi-tabs__mime",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json-schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details mime","data-collapsed":!1,open:!0,children:[(0,i.jsxs)("summary",{style:{},className:"openapi-markdown__details-summary-mime",children:[(0,i.jsx)("h3",{className:"openapi-markdown__details-summary-header-body",children:(0,i.jsx)(s.p,{children:"Body"})}),(0,i.jsx)("strong",{className:"openapi-schema__required",children:(0,i.jsx)(s.p,{children:"required"})})]}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:"1rem",marginBottom:"1rem"},children:(0,i.jsx)(s.p,{children:"Add Message Request"})})}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(x(),{collapsible:!0,className:"schemaItem",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details",children:[(0,i.jsx)("summary",{style:{},children:(0,i.jsxs)("span",{className:"openapi-schema__container",children:[(0,i.jsx)("strong",{className:"openapi-schema__property",children:(0,i.jsx)(s.p,{children:"message"})}),(0,i.jsx)("span",{className:"openapi-schema__name",children:(0,i.jsx)(s.p,{children:"object"})})]})}),(0,i.jsxs)("div",{style:{marginLeft:"1rem"},children:[(0,i.jsx)(x(),{collapsible:!1,name:"assetIds",required:!1,schemaName:"string[]",qualifierMessage:void 0,schema:{description:"AssetIds defines the attachments the message has.",items:{type:"string"},type:"array"}}),(0,i.jsx)(x(),{collapsible:!1,name:"content",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:'Content of the message eg. "Hi, what\'s up?"',type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"createdAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"id",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"threadId",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:"ThreadId of the message.",type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"updatedAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"userId",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:"UserId is the id of the user who wrote the message.\nFor AI messages this field is empty.",type:"string"}})]})]})})})]})})}),"\n",(0,i.jsx)("div",{children:(0,i.jsx)("div",{children:(0,i.jsxs)(r(),{label:void 0,id:void 0,children:[(0,i.jsxs)(f.default,{label:"200",value:"200",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Message successfully added"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsxs)(j(),{className:"openapi-tabs__schema",children:[(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(x(),{name:"property name*",required:!1,schemaName:"any",qualifierMessage:void 0,schema:{additionalProperties:!0,type:"object"},collapsible:!1,discriminator:!1})})]})}),(0,i.jsx)(f.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,i.jsx)(p(),{responseExample:"{}",language:"json"})})]})})})})]}),(0,i.jsxs)(f.default,{label:"400",value:"400",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Invalid JSON"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(j(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,i.jsxs)(f.default,{label:"401",value:"401",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Unauthorized"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(j(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,i.jsxs)(f.default,{label:"500",value:"500",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Internal Server Error"})}),(0,i.jsx)("div",{children:(0,i.jsx)(o(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(j(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]})]})})})]})}function q(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,i.jsx)(s,{...e,children:(0,i.jsx)(A,{...e})}):A(e)}}}]);