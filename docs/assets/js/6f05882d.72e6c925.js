"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5480],{2397:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>k,contentTitle:()=>N,default:()=>q,frontMatter:()=>b,metadata:()=>_,toc:()=>I});var i=a(74848),n=a(28453),l=a(91366),r=a.n(l),t=(a(6050),a(57742)),d=a.n(t),c=(a(67792),a(27362)),m=a.n(c),o=a(36683),p=a.n(o),h=a(81124),j=a.n(h),u=a(60674),x=a.n(u),g=a(23397),y=a.n(g),v=(a(26651),a(51107)),f=(a(77675),a(19365));const b={id:"add-message",title:"Add Message",description:"Add a new message to a specific thread.",sidebar_label:"Add Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVktv4zYQ/ivEXNoCipRNukChS+ttd7cu2iaIvYci8WEijiVuKZJLUva6hv97MXrEii0EDaqTOBzO45uZj9yDpFB45aKyBnKYSSlQGNqKmkLAkkS0AkVwVKi1KkSsPKFMIQGHHmuK5APk9/sTM8tWTcx/gQQUCxzGChIwWBPk0FmZS0jA05dGeZKQR99QAqGoqEbI9xB3jnVD9MqUcDisOmUK8Z2VO9YorIlkIv+ic1oVyO6zz4Fj2I9MOW8d+ago8KrP7HwDQ6A4l+3/CSz9jpC0VoaCiBUJjBGLqiYTu/WAWIWBAVKR6jCRSDII0Hvc8XqUx3O3P3cbwq6fOaAyFQ/wq0rEtsL4TRCN+/EBIDn3VHjCSHIWJ+NQcjq8oTpn8Sz7nZOA0inXjZMvuG4C+SkPn1q5UB2i6skT64ttZcXW20jPnD+YD9aL2XwQ8FEVxFqRbg1R7eJuIsRjIezjZyripOSQTAzIH30d7rp2POviAwuCsyZ0fXV1eXme6WAkNEVBIawbrXcCpSSeiv/c2Sil4i3Ut6NW7kbpLJdDAt9PhTI3G9RKit8WN3++xvkpoJ2DNxNVNdjEynr1z+uym3LwdjqDSN6gFgvyG/LivffW/z9P3AxYMrlBUWGEVQI1xcpKZjPbVr0ltRwy3r8I5DeqoKybnmw/TNEhG/gmgdBG1xFm4zWfzbQtUFc2xPztD1fXb4B5LlDReBV3C46v66F3GFQxa9jhU7RVjA56yuT1I+u0bcwJ3h3J8v1XrJ2mE/I70t39kPhqBNkge8YiRyGTx3F15IyjbMQAI2E/+OMxVGZtBz7Hoq0W1agYn4Cawk9BmbLRGL01aWHr40Uyu52LReOc9VyPDlJGJc+y0DjyTmNcW1+nqDI4m+W7xgg0UkjakLZOBNLrCy4FSeYTdC6k4i/beOG8LT3WNT5qEspcVLYJJD7eLlOxrEh8UJ4eMZBYW99yE5/ueFGrgkxo8R5i/nj7u9hcp5fPIg55lm2327Q0TWp9mfXnQoal0xfX6WVaxVq39Ey+DjfrRddvo4S3WJbkU2WzViXjsqjIZYfFEUBIgHuwQ+AyvWKT3M81mnGQR5qDE9hG9+4rngt9y0b6GjOnUZn2FuD09/0cdXM2zNFTT0EC+ei9MPTvKgGuFB/b7xn7T14fDiz+0pDfQX6/SmCDXnHJuheKCvwvIV+jDvRCVt/e9XT+nRg/ZCZTGO5ys2NgUTe8ggT+pt34oXNYHRKoCCX5Nphuu7/fL5Zs5Hj8jKoOyXBiVhTk4ou6Y6K6vVksIYHH/sFUW8lnPG75zsJtF6ptIWhZppXtQaMpm5YloLPJ37/Mg35k",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},N=void 0,_={id:"singulatron/add-message",title:"Add Message",description:"Add a new message to a specific thread.",source:"@site/docs/singulatron/add-message.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/add-message",permalink:"/docs/singulatron/add-message",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"add-message",title:"Add Message",description:"Add a new message to a specific thread.",sidebar_label:"Add Message",hide_title:!0,hide_table_of_contents:!0,api:"eJylVktv4zYQ/ivEXNoCipRNukChS+ttd7cu2iaIvYci8WEijiVuKZJLUva6hv97MXrEii0EDaqTOBzO45uZj9yDpFB45aKyBnKYSSlQGNqKmkLAkkS0AkVwVKi1KkSsPKFMIQGHHmuK5APk9/sTM8tWTcx/gQQUCxzGChIwWBPk0FmZS0jA05dGeZKQR99QAqGoqEbI9xB3jnVD9MqUcDisOmUK8Z2VO9YorIlkIv+ic1oVyO6zz4Fj2I9MOW8d+ago8KrP7HwDQ6A4l+3/CSz9jpC0VoaCiBUJjBGLqiYTu/WAWIWBAVKR6jCRSDII0Hvc8XqUx3O3P3cbwq6fOaAyFQ/wq0rEtsL4TRCN+/EBIDn3VHjCSHIWJ+NQcjq8oTpn8Sz7nZOA0inXjZMvuG4C+SkPn1q5UB2i6skT64ttZcXW20jPnD+YD9aL2XwQ8FEVxFqRbg1R7eJuIsRjIezjZyripOSQTAzIH30d7rp2POviAwuCsyZ0fXV1eXme6WAkNEVBIawbrXcCpSSeiv/c2Sil4i3Ut6NW7kbpLJdDAt9PhTI3G9RKit8WN3++xvkpoJ2DNxNVNdjEynr1z+uym3LwdjqDSN6gFgvyG/LivffW/z9P3AxYMrlBUWGEVQI1xcpKZjPbVr0ltRwy3r8I5DeqoKybnmw/TNEhG/gmgdBG1xFm4zWfzbQtUFc2xPztD1fXb4B5LlDReBV3C46v66F3GFQxa9jhU7RVjA56yuT1I+u0bcwJ3h3J8v1XrJ2mE/I70t39kPhqBNkge8YiRyGTx3F15IyjbMQAI2E/+OMxVGZtBz7Hoq0W1agYn4Cawk9BmbLRGL01aWHr40Uyu52LReOc9VyPDlJGJc+y0DjyTmNcW1+nqDI4m+W7xgg0UkjakLZOBNLrCy4FSeYTdC6k4i/beOG8LT3WNT5qEspcVLYJJD7eLlOxrEh8UJ4eMZBYW99yE5/ueFGrgkxo8R5i/nj7u9hcp5fPIg55lm2327Q0TWp9mfXnQoal0xfX6WVaxVq39Ey+DjfrRddvo4S3WJbkU2WzViXjsqjIZYfFEUBIgHuwQ+AyvWKT3M81mnGQR5qDE9hG9+4rngt9y0b6GjOnUZn2FuD09/0cdXM2zNFTT0EC+ei9MPTvKgGuFB/b7xn7T14fDiz+0pDfQX6/SmCDXnHJuheKCvwvIV+jDvRCVt/e9XT+nRg/ZCZTGO5ys2NgUTe8ggT+pt34oXNYHRKoCCX5Nphuu7/fL5Zs5Hj8jKoOyXBiVhTk4ou6Y6K6vVksIYHH/sFUW8lnPG75zsJtF6ptIWhZppXtQaMpm5YloLPJ37/Mg35k",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Update Thread",permalink:"/docs/singulatron/update-thread"},next:{title:"Get Messages",permalink:"/docs/singulatron/get-messages"}},k={},I=[];function T(e){const s={p:"p",...(0,n.R)(),...e.components},{Details:a}=s;return a||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(v.default,{as:"h1",className:"openapi__heading",children:"Add Message"}),"\n",(0,i.jsx)(d(),{method:"post",path:"/chat-service/thread/{threadId}/message"}),"\n",(0,i.jsx)(s.p,{children:"Add a new message to a specific thread."}),"\n",(0,i.jsx)(v.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsxs)(a,{style:{marginBottom:"1rem"},className:"openapi-markdown__details","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},children:(0,i.jsx)("h3",{className:"openapi-markdown__details-summary-header-params",children:(0,i.jsx)(s.p,{children:"Path Parameters"})})}),(0,i.jsx)("div",{children:(0,i.jsx)("ul",{children:(0,i.jsx)(p(),{className:"paramsItem",param:{description:"Thread ID",in:"path",name:"threadId",required:!0,schema:{type:"string"}}})})})]}),"\n",(0,i.jsx)(m(),{className:"openapi-tabs__mime",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json-schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details mime","data-collapsed":!1,open:!0,children:[(0,i.jsxs)("summary",{style:{},className:"openapi-markdown__details-summary-mime",children:[(0,i.jsx)("h3",{className:"openapi-markdown__details-summary-header-body",children:(0,i.jsx)(s.p,{children:"Body"})}),(0,i.jsx)("strong",{className:"openapi-schema__required",children:(0,i.jsx)(s.p,{children:"required"})})]}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:"1rem",marginBottom:"1rem"},children:(0,i.jsx)(s.p,{children:"Add Message Request"})})}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(x(),{collapsible:!0,className:"schemaItem",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details",children:[(0,i.jsx)("summary",{style:{},children:(0,i.jsxs)("span",{className:"openapi-schema__container",children:[(0,i.jsx)("strong",{className:"openapi-schema__property",children:(0,i.jsx)(s.p,{children:"message"})}),(0,i.jsx)("span",{className:"openapi-schema__name",children:(0,i.jsx)(s.p,{children:"object"})})]})}),(0,i.jsxs)("div",{style:{marginLeft:"1rem"},children:[(0,i.jsx)(x(),{collapsible:!1,name:"assetIds",required:!1,schemaName:"string[]",qualifierMessage:void 0,schema:{description:"AssetIds defines the attachments the message has.",items:{type:"string"},type:"array"}}),(0,i.jsx)(x(),{collapsible:!1,name:"content",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:'Content of the message eg. "Hi, what\'s up?"',type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"createdAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"id",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"threadId",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:"ThreadId of the message.",type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"updatedAt",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,i.jsx)(x(),{collapsible:!1,name:"userId",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{description:"UserId is the id of the user who wrote the message.\nFor AI messages this field is empty.",type:"string"}})]})]})})})]})})}),"\n",(0,i.jsx)("div",{children:(0,i.jsx)("div",{children:(0,i.jsxs)(r(),{label:void 0,id:void 0,children:[(0,i.jsxs)(f.default,{label:"200",value:"200",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Message successfully added"})}),(0,i.jsx)("div",{children:(0,i.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsxs)(y(),{className:"openapi-tabs__schema",children:[(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)(x(),{name:"property name*",required:!1,schemaName:"any",qualifierMessage:void 0,schema:{additionalProperties:!0,type:"object"},collapsible:!1,discriminator:!1})})]})}),(0,i.jsx)(f.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,i.jsx)(j(),{responseExample:"{}",language:"json"})})]})})})})]}),(0,i.jsxs)(f.default,{label:"400",value:"400",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Invalid JSON"})}),(0,i.jsx)("div",{children:(0,i.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(y(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,i.jsxs)(f.default,{label:"401",value:"401",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Unauthorized"})}),(0,i.jsx)("div",{children:(0,i.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(y(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,i.jsxs)(f.default,{label:"500",value:"500",children:[(0,i.jsx)("div",{children:(0,i.jsx)(s.p,{children:"Internal Server Error"})}),(0,i.jsx)("div",{children:(0,i.jsx)(m(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,i.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,i.jsx)(y(),{className:"openapi-tabs__schema",children:(0,i.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,i.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,i.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,i.jsx)("strong",{children:(0,i.jsx)(s.p,{children:"Schema"})})}),(0,i.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,i.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,i.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,i.jsx)(s.p,{children:"string"})})})]})})})})})})]})]})})})]})}function q(e={}){const{wrapper:s}={...(0,n.R)(),...e.components};return s?(0,i.jsx)(s,{...e,children:(0,i.jsx)(T,{...e})}):T(e)}}}]);