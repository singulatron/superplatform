"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1561],{59473:(n,e,i)=>{i.r(e),i.d(e,{assets:()=>t,contentTitle:()=>a,default:()=>u,frontMatter:()=>s,metadata:()=>r,toc:()=>d});var l=i(74848),o=i(28453);const s={sidebar_position:1,tags:["run","deploy","local"]},a="Running Locally",r={id:"running/locally",title:"Running Locally",description:"The easiest way to run Singulatron is to use Docker Compose.",source:"@site/docs/running/locally.md",sourceDirName:"running",slug:"/running/locally",permalink:"/docs/running/locally",draft:!1,unlisted:!1,editUrl:"https://github.com/singulatron/singulatron/tree/main/docs-source/docs/running/locally.md",tags:[{inline:!0,label:"run",permalink:"/docs/tags/run"},{inline:!0,label:"deploy",permalink:"/docs/tags/deploy"},{inline:!0,label:"local",permalink:"/docs/tags/local"}],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1,tags:["run","deploy","local"]},sidebar:"tutorialSidebar",previous:{title:"Running",permalink:"/docs/category/running"},next:{title:"Docker Compose",permalink:"/docs/running/docker-compose"}},t={},d=[{value:"Docker Compose",id:"docker-compose",level:2},{value:"Once it&#39;s running",id:"once-its-running",level:3},{value:"Natively (Go &amp; Angular)",id:"natively-go--angular",level:2},{value:"Backend",id:"backend",level:2},{value:"Frontend",id:"frontend",level:2},{value:"Once it&#39;s running",id:"once-its-running-1",level:3},{value:"Administration",id:"administration",level:2},{value:"Local files",id:"local-files",level:3},{value:"Config file",id:"config-file",level:4},{value:"Download.json",id:"downloadjson",level:4},{value:"Data files",id:"data-files",level:4}];function c(n){const e={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",h4:"h4",p:"p",pre:"pre",...(0,o.R)(),...n.components};return(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)(e.h1,{id:"running-locally",children:"Running Locally"}),"\n",(0,l.jsxs)(e.p,{children:["The easiest way to run Singulatron is to use ",(0,l.jsx)(e.a,{href:"https://docs.docker.com/compose/install/",children:"Docker Compose"}),"."]}),"\n",(0,l.jsx)(e.h2,{id:"docker-compose",children:"Docker Compose"}),"\n",(0,l.jsxs)(e.p,{children:["The easiest way to run this is to clone ",(0,l.jsx)(e.a,{href:"https://github.com/singulatron/singulatron",children:"the repo"}),", step into the repo root and run:"]}),"\n",(0,l.jsx)(e.pre,{children:(0,l.jsx)(e.code,{className:"language-sh",children:"git clone git@github.com:singulatron/singulatron.git\ncd singulatron\ndocker compose up\n# or use the -d flag to run it in the background\n# docker compose up -d\n"})}),"\n",(0,l.jsxs)(e.p,{children:["The ",(0,l.jsx)(e.code,{children:"docker-compose-yaml"})," in the root folder is designed to build and run the current code. For a more production ready Docker Compose file see the ",(0,l.jsx)(e.a,{href:"./docker-compose/",children:"Docker Compose page"}),"."]}),"\n",(0,l.jsx)(e.h3,{id:"once-its-running",children:"Once it's running"}),"\n",(0,l.jsxs)(e.p,{children:["After the containers successfully start, you can go to ",(0,l.jsx)(e.a,{href:"http://127.0.0.1:3901",children:"http://127.0.0.1:3901"})," and log in with the ",(0,l.jsx)(e.a,{href:"/docs/running/using#default-credentials",children:"Default Credentials"}),"."]}),"\n",(0,l.jsx)(e.h2,{id:"natively-go--angular",children:"Natively (Go & Angular)"}),"\n",(0,l.jsx)(e.p,{children:"If you have both Go and Angular installed on your computer, the easiest way to dip your feet into Singulatron is to run things locally."}),"\n",(0,l.jsx)(e.h2,{id:"backend",children:"Backend"}),"\n",(0,l.jsx)(e.pre,{children:(0,l.jsx)(e.code,{className:"language-bash",children:"cd localtron;\ngo run main.go\n"})}),"\n",(0,l.jsx)(e.h2,{id:"frontend",children:"Frontend"}),"\n",(0,l.jsx)(e.pre,{children:(0,l.jsx)(e.code,{className:"language-bash",children:"cd desktop/workspaces/angular-app/;\nnpm run start\n"})}),"\n",(0,l.jsx)(e.h3,{id:"once-its-running-1",children:"Once it's running"}),"\n",(0,l.jsxs)(e.p,{children:["After the both the backend and frontend starts, you can go to ",(0,l.jsx)(e.a,{href:"http://127.0.0.1:4200",children:"http://127.0.0.1:4200"})," and log in with the ",(0,l.jsx)(e.a,{href:"/docs/running/using#default-credentials",children:"Default Credentials"}),"."]}),"\n",(0,l.jsx)(e.h2,{id:"administration",children:"Administration"}),"\n",(0,l.jsx)(e.h3,{id:"local-files",children:"Local files"}),"\n",(0,l.jsxs)(e.p,{children:["By default Singulatron uses the folder ",(0,l.jsx)(e.code,{children:"~/.singulatron"})," on your machine for config, file downloads and for the local database."]}),"\n",(0,l.jsx)(e.h4,{id:"config-file",children:"Config file"}),"\n",(0,l.jsx)(e.pre,{children:(0,l.jsx)(e.code,{className:"language-bash",children:"cat ~/.singulatron/config.yaml\n"})}),"\n",(0,l.jsx)(e.h4,{id:"downloadjson",children:"Download.json"}),"\n",(0,l.jsx)(e.p,{children:"This file contains all the local downloads on a node. Losing is file is not a big deal as downloaded files are detected even if this file or the entry in this file is missing."}),"\n",(0,l.jsx)(e.pre,{children:(0,l.jsx)(e.code,{className:"language-bash",children:"~/.singulatron/downloads.json\n"})}),"\n",(0,l.jsx)(e.h4,{id:"data-files",children:"Data files"}),"\n",(0,l.jsx)(e.p,{children:"By default Singulatron uses local gzipped json files to store database entries. Data access across Singulatron is interface based so the this implementation can be easily swapped out for PostgreSQL and other database backends."}),"\n",(0,l.jsx)(e.p,{children:"The files are located at"}),"\n",(0,l.jsx)(e.pre,{children:(0,l.jsx)(e.code,{className:"language-bash",children:"ls ~/.singulatron/data\n"})}),"\n",(0,l.jsx)(e.p,{children:"If you want to view the contents of a file:"}),"\n",(0,l.jsx)(e.pre,{children:(0,l.jsx)(e.code,{className:"language-bash",children:"cat ~/.singulatron/data/users.zip | gzip -dc\n\n# or if you jave jq installed\ncat ~/.singulatron/data/users.zip | gzip -dc | jq\n"})})]})}function u(n={}){const{wrapper:e}={...(0,o.R)(),...n.components};return e?(0,l.jsx)(e,{...n,children:(0,l.jsx)(c,{...n})}):c(n)}}}]);