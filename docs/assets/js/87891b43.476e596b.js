"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4403],{3535:(e,s,a)=>{a.r(s),a.d(s,{assets:()=>b,contentTitle:()=>N,default:()=>L,frontMatter:()=>y,metadata:()=>_,toc:()=>v});var n=a(74848),i=a(28453),l=a(91366),r=a.n(l),t=(a(6050),a(57742)),o=a.n(t),d=(a(67792),a(27362)),c=a.n(d),m=(a(36683),a(81124)),p=a.n(m),h=a(60674),j=a.n(h),g=a(23397),u=a.n(g),x=(a(26651),a(51107)),f=(a(77675),a(19365));const y={id:"get-config",title:"Get Config",description:"Fetch the current configuration from the server",sidebar_label:"Get Config",hide_title:!0,hide_table_of_contents:!0,api:"eJylVttu20YQ/ZXFPAY0qdgNUPCpShobQt3GsGSgha2HETkiN1nubvZiRRH478WQ1MWSGrjtC0nMzvXMnFluoCRfOGmDNBpyuKZQ1CLUJIroHOkgCqOXsooOWUMsnWm6Y0/umRwk4OhrJB/em3IN+QYKowPpwJ9orZJFZ5d99ux+A76oqUH+CmtLkINZfKYiQNu2yVEqNxTEhy64uO9jDNGkoxLy4CK1LPDWaE+efV6ORvx66efD2UIcBSfpmUrhY1GQ98uo1BqS1xdgnbHkguxD995P5WjtqVCZqpK6+lV6XCguZgfHwhhFqIHhOEIogVI6KoJx6wMDH5zUVXdqVloZLE/DbU+ujSrJnTE+F0z6+6iDbGiifUA1pPkS2jdiVksvlgorUTr5TF6gaKQ2Tjz8KZaEITrKn/RTEG+EXHaDEz05UaMX2gQht767I9cHFCsSvjYrgVpM/pjOxre3g4tFDMHohN8n7lA5wnL9Y5eDn10AMZ2N72c7rd5/2mtBctKUBBpTkjrT/X7GfufTSfkqhF8lYdlPo7enyD9ojKE2Tn6n8t8M7XFaHODdOdZMdCCnUYlpR3Tx0Tnj/l+kNgFPRXQyrCF/3MB7QkduHEMN+eO8nScQsPKQP8JAew4tC4J5Ag2F2pSQgzXdHrDIVpD1tMsqYmG/k3znPDrF55kyBara+JC/+/ny6i1wmG0WU062b+BhLsdQzNaWxNOg8gRiaZQyKyrFYi1QeIsFCdSlCOYLaYFFv6H2q/KBR9QPxYhbU0ktSJfWSB1SSEBykJqw7PapxoZxGw/t7eDdjyJa+Rutu1lhyO/3y/fjN2ysIsg3TF69NNttjEXXLWpQMiQeFflfvNRVVBic0WlhmoO4dxMxjdYax4j2KNYh2DzLfLTkrMKwNK5JUWZwsrPvo+6gKOmZlLHCk1peMPpUivFEoLU+FX+Z6IR1pnLYNLwAhdQXtYmexM3dLBWzmsS1dLRAT2JpXAciW1fEcClZkPZc6C7nm7tb8XyVjl5k7PMsW61WaaVjalyVDXY+w8qqi6t0lNahUVxDINf4T8vtuO0LXmFVkUulyTqVjPsgA4MM0z2AkACPXY/AKL1klzymDeqDJPfXGRyhdnBp/of7dxiMQN9CZhXKbk11KGwGljxuL6cEmCfzBLgjLN9sGOMHp9qWxV8j8d3yOE/gGZ3k1nTMTLbTycT6Qmu+VPuML5gbXD+q2I/n0Upok63FuCjIhh/qHvL87tN0Bgksht8K3ruQg8MV/wTgCnKABEyHYEfgTrYBhbqKWLFu75OZggOt9xzilJLtB1e1PdLrgwyPOdgXwk8u66zJZtMztG13+v3RP1rsiN9rc//mbdv+DTg2U+A=",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},N=void 0,_={id:"singulatron/get-config",title:"Get Config",description:"Fetch the current configuration from the server",source:"@site/docs/singulatron/get-config.api.mdx",sourceDirName:"singulatron",slug:"/singulatron/get-config",permalink:"/docs/singulatron/get-config",draft:!1,unlisted:!1,editUrl:null,tags:[],version:"current",frontMatter:{id:"get-config",title:"Get Config",description:"Fetch the current configuration from the server",sidebar_label:"Get Config",hide_title:!0,hide_table_of_contents:!0,api:"eJylVttu20YQ/ZXFPAY0qdgNUPCpShobQt3GsGSgha2HETkiN1nubvZiRRH478WQ1MWSGrjtC0nMzvXMnFluoCRfOGmDNBpyuKZQ1CLUJIroHOkgCqOXsooOWUMsnWm6Y0/umRwk4OhrJB/em3IN+QYKowPpwJ9orZJFZ5d99ux+A76oqUH+CmtLkINZfKYiQNu2yVEqNxTEhy64uO9jDNGkoxLy4CK1LPDWaE+efV6ORvx66efD2UIcBSfpmUrhY1GQ98uo1BqS1xdgnbHkguxD995P5WjtqVCZqpK6+lV6XCguZgfHwhhFqIHhOEIogVI6KoJx6wMDH5zUVXdqVloZLE/DbU+ujSrJnTE+F0z6+6iDbGiifUA1pPkS2jdiVksvlgorUTr5TF6gaKQ2Tjz8KZaEITrKn/RTEG+EXHaDEz05UaMX2gQht767I9cHFCsSvjYrgVpM/pjOxre3g4tFDMHohN8n7lA5wnL9Y5eDn10AMZ2N72c7rd5/2mtBctKUBBpTkjrT/X7GfufTSfkqhF8lYdlPo7enyD9ojKE2Tn6n8t8M7XFaHODdOdZMdCCnUYlpR3Tx0Tnj/l+kNgFPRXQyrCF/3MB7QkduHEMN+eO8nScQsPKQP8JAew4tC4J5Ag2F2pSQgzXdHrDIVpD1tMsqYmG/k3znPDrF55kyBara+JC/+/ny6i1wmG0WU062b+BhLsdQzNaWxNOg8gRiaZQyKyrFYi1QeIsFCdSlCOYLaYFFv6H2q/KBR9QPxYhbU0ktSJfWSB1SSEBykJqw7PapxoZxGw/t7eDdjyJa+Rutu1lhyO/3y/fjN2ysIsg3TF69NNttjEXXLWpQMiQeFflfvNRVVBic0WlhmoO4dxMxjdYax4j2KNYh2DzLfLTkrMKwNK5JUWZwsrPvo+6gKOmZlLHCk1peMPpUivFEoLU+FX+Z6IR1pnLYNLwAhdQXtYmexM3dLBWzmsS1dLRAT2JpXAciW1fEcClZkPZc6C7nm7tb8XyVjl5k7PMsW61WaaVjalyVDXY+w8qqi6t0lNahUVxDINf4T8vtuO0LXmFVkUulyTqVjPsgA4MM0z2AkACPXY/AKL1klzymDeqDJPfXGRyhdnBp/of7dxiMQN9CZhXKbk11KGwGljxuL6cEmCfzBLgjLN9sGOMHp9qWxV8j8d3yOE/gGZ3k1nTMTLbTycT6Qmu+VPuML5gbXD+q2I/n0Upok63FuCjIhh/qHvL87tN0Bgksht8K3ruQg8MV/wTgCnKABEyHYEfgTrYBhbqKWLFu75OZggOt9xzilJLtB1e1PdLrgwyPOdgXwk8u66zJZtMztG13+v3RP1rsiN9rc//mbdv+DTg2U+A=",sidebar_class_name:"post api-method",info_path:"docs/singulatron/singulatron",custom_edit_url:null},sidebar:"openApiSidebar",previous:{title:"Get Threads",permalink:"/docs/singulatron/get-threads"},next:{title:"Save Config",permalink:"/docs/singulatron/save-config"}},b={},v=[];function k(e){const s={p:"p",...(0,i.R)(),...e.components},{Details:a}=s;return a||function(e,s){throw new Error("Expected "+(s?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(x.default,{as:"h1",className:"openapi__heading",children:"Get Config"}),"\n",(0,n.jsx)(o(),{method:"post",path:"/config/get"}),"\n",(0,n.jsx)(s.p,{children:"Fetch the current configuration from the server"}),"\n",(0,n.jsx)(x.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(c(),{className:"openapi-tabs__mime",children:(0,n.jsx)(f.default,{label:"application/json",value:"application/json-schema",children:(0,n.jsxs)(a,{style:{},className:"openapi-markdown__details mime","data-collapsed":!1,open:!0,children:[(0,n.jsxs)("summary",{style:{},className:"openapi-markdown__details-summary-mime",children:[(0,n.jsx)("h3",{className:"openapi-markdown__details-summary-header-body",children:(0,n.jsx)(s.p,{children:"Body"})}),(0,n.jsx)("strong",{className:"openapi-schema__required",children:(0,n.jsx)(s.p,{children:"required"})})]}),(0,n.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"},children:(0,n.jsx)("div",{style:{marginTop:"1rem",marginBottom:"1rem"},children:(0,n.jsx)(s.p,{children:"Get Config Request"})})}),(0,n.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,n.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,n.jsx)(s.p,{children:"object"})})})]})})}),"\n",(0,n.jsx)("div",{children:(0,n.jsx)("div",{children:(0,n.jsxs)(r(),{label:void 0,id:void 0,children:[(0,n.jsxs)(f.default,{label:"200",value:"200",children:[(0,n.jsx)("div",{children:(0,n.jsx)(s.p,{children:"Current configuration retrieved successfully"})}),(0,n.jsx)("div",{children:(0,n.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,n.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,n.jsxs)(u(),{className:"openapi-tabs__schema",children:[(0,n.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,n.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,n.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,n.jsx)("strong",{children:(0,n.jsx)(s.p,{children:"Schema"})})}),(0,n.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,n.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,n.jsx)(j(),{collapsible:!0,className:"schemaItem",children:(0,n.jsxs)(a,{style:{},className:"openapi-markdown__details",children:[(0,n.jsx)("summary",{style:{},children:(0,n.jsxs)("span",{className:"openapi-schema__container",children:[(0,n.jsx)("strong",{className:"openapi-schema__property",children:(0,n.jsx)(s.p,{children:"config"})}),(0,n.jsx)("span",{className:"openapi-schema__name",children:(0,n.jsx)(s.p,{children:"object"})})]})}),(0,n.jsxs)("div",{style:{marginLeft:"1rem"},children:[(0,n.jsx)(j(),{collapsible:!0,className:"schemaItem",children:(0,n.jsxs)(a,{style:{},className:"openapi-markdown__details",children:[(0,n.jsx)("summary",{style:{},children:(0,n.jsxs)("span",{className:"openapi-schema__container",children:[(0,n.jsx)("strong",{className:"openapi-schema__property",children:(0,n.jsx)(s.p,{children:"app"})}),(0,n.jsx)("span",{className:"openapi-schema__name",children:(0,n.jsx)(s.p,{children:"object"})})]})}),(0,n.jsx)("div",{style:{marginLeft:"1rem"},children:(0,n.jsx)(j(),{collapsible:!1,name:"loggingDisabled",required:!1,schemaName:"boolean",qualifierMessage:void 0,schema:{type:"boolean"}})})]})}),(0,n.jsx)(j(),{collapsible:!1,name:"directory",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}}),(0,n.jsx)(j(),{collapsible:!0,className:"schemaItem",children:(0,n.jsxs)(a,{style:{},className:"openapi-markdown__details",children:[(0,n.jsx)("summary",{style:{},children:(0,n.jsxs)("span",{className:"openapi-schema__container",children:[(0,n.jsx)("strong",{className:"openapi-schema__property",children:(0,n.jsx)(s.p,{children:"download"})}),(0,n.jsx)("span",{className:"openapi-schema__name",children:(0,n.jsx)(s.p,{children:"object"})})]})}),(0,n.jsx)("div",{style:{marginLeft:"1rem"},children:(0,n.jsx)(j(),{collapsible:!1,name:"downloadFolder",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})}),(0,n.jsx)(j(),{collapsible:!1,name:"isRuntimeInstalled",required:!1,schemaName:"boolean",qualifierMessage:void 0,schema:{description:"* This flag drives a minor UX feature:\n\t * if the user has not installed the runtime we show an INSTALL\n\t * button, but if the user has already installed the runtime we show\n\t * we show a START runtime button.\n\t *",type:"boolean"}}),(0,n.jsx)(j(),{collapsible:!0,className:"schemaItem",children:(0,n.jsxs)(a,{style:{},className:"openapi-markdown__details",children:[(0,n.jsx)("summary",{style:{},children:(0,n.jsxs)("span",{className:"openapi-schema__container",children:[(0,n.jsx)("strong",{className:"openapi-schema__property",children:(0,n.jsx)(s.p,{children:"model"})}),(0,n.jsx)("span",{className:"openapi-schema__name",children:(0,n.jsx)(s.p,{children:"object"})})]})}),(0,n.jsx)("div",{style:{marginLeft:"1rem"},children:(0,n.jsx)(j(),{collapsible:!1,name:"currentModelId",required:!1,schemaName:"string",qualifierMessage:void 0,schema:{type:"string"}})})]})})]})]})})})]})}),(0,n.jsx)(f.default,{label:"Example (from schema)",value:"Example (from schema)",children:(0,n.jsx)(p(),{responseExample:'{\n  "config": {\n    "app": {\n      "loggingDisabled": true\n    },\n    "directory": "string",\n    "download": {\n      "downloadFolder": "string"\n    },\n    "isRuntimeInstalled": true,\n    "model": {\n      "currentModelId": "string"\n    }\n  }\n}',language:"json"})})]})})})})]}),(0,n.jsxs)(f.default,{label:"401",value:"401",children:[(0,n.jsx)("div",{children:(0,n.jsx)(s.p,{children:"Unauthorized"})}),(0,n.jsx)("div",{children:(0,n.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,n.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,n.jsx)(u(),{className:"openapi-tabs__schema",children:(0,n.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,n.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,n.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,n.jsx)("strong",{children:(0,n.jsx)(s.p,{children:"Schema"})})}),(0,n.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,n.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,n.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,n.jsx)(s.p,{children:"string"})})})]})})})})})})]}),(0,n.jsxs)(f.default,{label:"500",value:"500",children:[(0,n.jsx)("div",{children:(0,n.jsx)(s.p,{children:"Internal Server Error"})}),(0,n.jsx)("div",{children:(0,n.jsx)(c(),{className:"openapi-tabs__mime",schemaType:"response",children:(0,n.jsx)(f.default,{label:"application/json",value:"application/json",children:(0,n.jsx)(u(),{className:"openapi-tabs__schema",children:(0,n.jsx)(f.default,{label:"Schema",value:"Schema",children:(0,n.jsxs)(a,{style:{},className:"openapi-markdown__details response","data-collapsed":!1,open:!0,children:[(0,n.jsx)("summary",{style:{},className:"openapi-markdown__details-summary-response",children:(0,n.jsx)("strong",{children:(0,n.jsx)(s.p,{children:"Schema"})})}),(0,n.jsx)("div",{style:{textAlign:"left",marginLeft:"1rem"}}),(0,n.jsx)("ul",{style:{marginLeft:"1rem"},children:(0,n.jsx)("div",{style:{marginTop:".5rem",marginBottom:".5rem"},children:(0,n.jsx)(s.p,{children:"string"})})})]})})})})})})]})]})})})]})}function L(e={}){const{wrapper:s}={...(0,i.R)(),...e.components};return s?(0,n.jsx)(s,{...e,children:(0,n.jsx)(k,{...e})}):k(e)}}}]);