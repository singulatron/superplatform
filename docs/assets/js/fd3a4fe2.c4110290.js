"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[5930],{34748:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>w,contentTitle:()=>y,default:()=>_,frontMatter:()=>b,metadata:()=>v,toc:()=>N});var l=a(74848),i=a(28453),n=a(91366),t=a.n(n),r=(a(6050),a(57742)),o=a.n(r),d=(a(67792),a(27362)),c=a.n(d),p=(a(36683),a(81124)),m=a.n(p),h=a(60674),u=a.n(h),x=a(23397),j=a.n(x),f=(a(26651),a(51107)),g=(a(77675),a(19365));const b={id:"list-downloads",title:"List Downloads",description:"Fetch a list of all download details.",sidebar_label:"List Downloads",hide_title:!0,hide_table_of_contents:!0,api:"eJy1VU1v20YQ/SuLOdOkYjdAwVMdtA6MGo0R2YfCFpAROSI3We5udpZSFYH/PZilFFGOfCiKXiRy9s3Hvpk33EFNXAXto3YWSrihWLUKldEclVspNEbVbmONw1rVFFEbzp/ts/1IX3sdiFVsSX06QC54XZWHl3KtafNJeQqdZtbO5pCB8xRQkt3WUIKk+X0PZ8ggEHtnmRjKHVzOZvJ3WuDdvrB64lU5G8lGAaP3RlcpQfGZxWMHXLXUoTz5IOmjHuMfQ5Q70JE6/hlToa3IGKrlJW49QQlL5wyhhSGDWofJAcegbZPs+9BUv9vGMdIeo22khoKAKAR33n2lDf2FHb16eI+xPX/YG3OjDc31NzqfVddnHT32/NotfXBNIJ5ew/bdcozHEWPPZ2P2wZyxD9nB4pafqYpwNGAIuD2HGMT2y+zNz/PwaLGPrQv6G9X/ZhReFiUJ3p4buFsbKVg0ak5hTUH9kbr2nzIJa1T1QcctlE87eEcYKFz30tOnxbDIIGLDUD7BQRspua4IFhl0FFsn4vGOI0jnxA+KqQaLqTw4Fc4pVWoJFIVxFZrWcSzf/np59QYk6aGmuZQ+Du20spfEPGw9qec95BnUyhnjNlSr5VahYo8VKbS1iu4LWYVV2he1WgXXpaXxyBQU7y+m7lyjrSJbe6dtlE2hJUlLWJPQbZMc4Hrf7EQ2HAfH6z9pnBxtVy4p19mIVeoPdajl2oyG+DfWtukNxuBsXrluEvv+Vs17710QWkem2hh9WRTcewreYFy50OWoiyTyEzo+9jZdt6Y1GecVk1ldCMNUq+tbhd5zrv52fVBJTth1uDSktL1oXc+k3t8/5OqhJXWjAy2RSa1cSESJd0NCidEVWU7CPtT8/v5Ora/y2UnFXBbFZrPJG9vnLjTF3o8LbLy5uMpneRs7k4RHoeMPq8N4HS+8waahkGtXJEghXOtoBDI/EggZyGiNDMzyy7QtHMcO7aTItLOna/6Eud1RSv//52c/L5H+iYU3qO1kT41CeoJpNMgmn4lFBtJQwex20qLHYIZBzF97CqLlRQZrDFo6m6ScHQZYtPeFttKxqiIvA7ZG04+z+2J7DFOR33+YP0AG2J8sfPRaomWHBwl/OLLbSfCXehlrkN8he8VltxvVNAw/8OPRqx4/RDqihdTFMAzfAaLf5Cw=",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},y=void 0,v={id:"singulatron/list-downloads",title:"List Downloads",description:"Fetch a list of all download details.",source:"@site/docs/singulatron/list-downloads.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/list-downloads",permalink:"/docs/singulatron/list-downloads",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"list-downloads",title:"List Downloads",description:"Fetch a list of all download details.",sidebar_label:"List Downloads",hide_title:!0,hide_table_of_contents:!0,api:"eJy1VU1v20YQ/SuLOdOkYjdAwVMdtA6MGo0R2YfCFpAROSI3We5udpZSFYH/PZilFFGOfCiKXiRy9s3Hvpk33EFNXAXto3YWSrihWLUKldEclVspNEbVbmONw1rVFFEbzp/ts/1IX3sdiFVsSX06QC54XZWHl3KtafNJeQqdZtbO5pCB8xRQkt3WUIKk+X0PZ8ggEHtnmRjKHVzOZvJ3WuDdvrB64lU5G8lGAaP3RlcpQfGZxWMHXLXUoTz5IOmjHuMfQ5Q70JE6/hlToa3IGKrlJW49QQlL5wyhhSGDWofJAcegbZPs+9BUv9vGMdIeo22khoKAKAR33n2lDf2FHb16eI+xPX/YG3OjDc31NzqfVddnHT32/NotfXBNIJ5ew/bdcozHEWPPZ2P2wZyxD9nB4pafqYpwNGAIuD2HGMT2y+zNz/PwaLGPrQv6G9X/ZhReFiUJ3p4buFsbKVg0ak5hTUH9kbr2nzIJa1T1QcctlE87eEcYKFz30tOnxbDIIGLDUD7BQRspua4IFhl0FFsn4vGOI0jnxA+KqQaLqTw4Fc4pVWoJFIVxFZrWcSzf/np59QYk6aGmuZQ+Du20spfEPGw9qec95BnUyhnjNlSr5VahYo8VKbS1iu4LWYVV2he1WgXXpaXxyBQU7y+m7lyjrSJbe6dtlE2hJUlLWJPQbZMc4Hrf7EQ2HAfH6z9pnBxtVy4p19mIVeoPdajl2oyG+DfWtukNxuBsXrluEvv+Vs17710QWkem2hh9WRTcewreYFy50OWoiyTyEzo+9jZdt6Y1GecVk1ldCMNUq+tbhd5zrv52fVBJTth1uDSktL1oXc+k3t8/5OqhJXWjAy2RSa1cSESJd0NCidEVWU7CPtT8/v5Ora/y2UnFXBbFZrPJG9vnLjTF3o8LbLy5uMpneRs7k4RHoeMPq8N4HS+8waahkGtXJEghXOtoBDI/EggZyGiNDMzyy7QtHMcO7aTItLOna/6Eud1RSv//52c/L5H+iYU3qO1kT41CeoJpNMgmn4lFBtJQwex20qLHYIZBzF97CqLlRQZrDFo6m6ScHQZYtPeFttKxqiIvA7ZG04+z+2J7DFOR33+YP0AG2J8sfPRaomWHBwl/OLLbSfCXehlrkN8he8VltxvVNAw/8OPRqx4/RDqihdTFMAzfAaLf5Cw=",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Pause a Download",permalink:"/docs/singulatron/pause"},next:{title:"Publish an Event",permalink:"/docs/singulatron/publish-an-event"}},w={},N=[];function F(e){const s={code:"code",p:"p",...(0,i.R)(),...e.components},{Details:a}=s;return a||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"List Downloads"}),"\n",(0,l.jsx)(o(),{method:"post",path:"/download-svc/downloads"}),"\n",(0,l.jsx)(s.p,{children:"Fetch a list of all download details."}),"\n",(0,l.jsxs)(s.p,{children:["Requires the ",(0,l.jsx)(s.code,{children:"download-svc:download:view"})," permission."]}),"\n",(0,l.jsx)("div",{children:(0,l.jsx)("div",{children:(0,l.jsxs)(t(),{label:void 0,id:void 0,children:[(0,l.jsxs)(g.default,{label:"200",value:"200",children:[(0,l.jsx)("div",{children:(0,l.jsx)(s.p,{children:"List of downloads"})}),(0,l.jsx)("div",{children:(0,l.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(g.default,{label:"application/json",value:"application/json",children:(0,l.jsxs)(j(),{className:"openapi-tabs__schema",children:[(0,l.jsx)(g.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(s.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)(u(),{collapsible:!0,className:"schemaItem",children:(0,l.jsxs)(a,{style:{},className:"openapi-markdown__details",children:[(0,l.jsx)("summary",{style:{},children:(0,l.jsxs)("span",{className:"openapi-schema__container",children:[(0,l.jsx)("strong",{className:"openapi-schema__property",children:(0,l.jsx)(s.p,{children:"downloads"})}),(0,l.jsx)("span",{className:"openapi-schema__name",children:(0,l.jsx)(s.p,{children:"object[]"})})]})}),(0,l.jsxs)("div",{style:{marginLeft:"1rem"},children:[(0,l.jsx)("li",{children:(0,l.jsx)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem",paddingBottom:".5rem"},children:(0,l.jsx)(s.p,{children:"Array ["})})}),(0,l.jsx)(u(),{collapsible:!1,name:"cancelled",required:!1,schemaName:"boolean",qualifierMessage:void 0,schema:{type:"boolean"}}),(0,l.jsx)(u(),{collapsible:!1,name:"dir",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,l.jsx)(u(),{collapsible:!1,name:"downloadedBytes",required:!1,schemaName:"integer",qualifierMessage:void 0,schema:{type:"integer"}}),(0,l.jsx)(u(),{collapsible:!1,name:"error",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,l.jsx)(u(),{collapsible:!1,name:"fileName",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,l.jsx)(u(),{collapsible:!1,name:"filePath",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,l.jsx)(u(),{collapsible:!1,name:"fullFileSize",required:!1,schemaName:"integer",qualifierMessage:void 0,schema:{type:"integer"}}),(0,l.jsx)(u(),{collapsible:!1,name:"id",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,l.jsx)(u(),{collapsible:!1,name:"paused",required:!1,schemaName:"boolean",qualifierMessage:void 0,schema:{type:"boolean"}}),(0,l.jsx)(u(),{collapsible:!1,name:"progress",required:!1,schemaName:"number",qualifierMessage:void 0,schema:{type:"number"}}),(0,l.jsx)(u(),{collapsible:!1,name:"status",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,l.jsx)(u(),{collapsible:!1,name:"url",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,l.jsx)("li",{children:(0,l.jsx)("div",{style:{fontSize:"var(--ifm-code-font-size)",opacity:"0.6",marginLeft:"-.5rem"},children:(0,l.jsx)(s.p,{children:"]"})})})]})]})})})]})}),(0,l.jsx)(g.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,l.jsx)(m(),{responseExample:'{\n  "downloads": [\n    {\n      "cancelled": true,\n      "dir": "string",\n      "downloadedBytes": 0,\n      "error": "string",\n      "fileName": "string",\n      "filePath": "string",\n      "fullFileSize": 0,\n      "id": "string",\n      "paused": true,\n      "progress": 0,\n      "status": "string",\n      "url": "string"\n    }\n  ]\n}',language:"json"})})]})})})})]}),(0,l.jsxs)(g.default,{label:"401",value:"401",children:[(0,l.jsx)("div",{children:(0,l.jsx)(s.p,{children:"Unauthorized"})}),(0,l.jsx)("div",{children:(0,l.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(g.default,{label:"application/json",value:"application/json",children:(0,l.jsx)(j(),{className:"openapi-tabs__schema",children:(0,l.jsx)(g.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(s.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,l.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,l.jsxs)(g.default,{label:"500",value:"500",children:[(0,l.jsx)("div",{children:(0,l.jsx)(s.p,{children:"Internal Server Error"})}),(0,l.jsx)("div",{children:(0,l.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,l.jsx)(g.default,{label:"application/json",value:"application/json",children:(0,l.jsx)(j(),{className:"openapi-tabs__schema",children:(0,l.jsx)(g.default,{label:"Schema",value:"Schema",children:(0,l.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,l.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,l.jsx)("strong",{children:(0,l.jsx)(s.p,{children:"Schema"})})}),(0,l.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,l.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,l.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,l.jsx)(s.p,{children:"string"})})})]})})})})})})]})]})})})]})}function _(e={}){const{wrapper:s}={...(0,i.R)(),...e.components};return s?(0,l.jsx)(s,{...e,children:(0,l.jsx)(F,{...e})}):F(e)}}}]);