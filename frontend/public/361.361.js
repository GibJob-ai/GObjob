(self.webpackChunksvelte_app=self.webpackChunksvelte_app||[]).push([[361],{361:(t,e,n)=>{"use strict";n.r(e),n.d(e,{default:()=>T});var s=n(379);function l(t,e,n){const s=t.slice();return s[11]=e[n],s[12]=e,s[13]=n,s}function a(t,e,n){const s=t.slice();return s[14]=e[n],s[15]=e,s[16]=n,s}function o(t){let e,n;return{c(){e=(0,s.bG)("div"),n=(0,s.fL)(t[2])},m(t,l){(0,s.$T)(t,e,l),(0,s.R3)(e,n)},p(t,e){4&e&&(0,s.rT)(n,t[2])},d(t){t&&(0,s.og)(e)}}}function i(t){let e,n=t[0].related,l=[];for(let e=0;e<n.length;e+=1)l[e]=r(a(t,n,e));return{c(){for(let t=0;t<l.length;t+=1)l[t].c();e=(0,s.cS)()},m(t,n){for(let e=0;e<l.length;e+=1)l[e].m(t,n);(0,s.$T)(t,e,n)},p(t,s){if(1&s){let o;for(n=t[0].related,o=0;o<n.length;o+=1){const i=a(t,n,o);l[o]?l[o].p(i,s):(l[o]=r(i),l[o].c(),l[o].m(e.parentNode,e))}for(;o<l.length;o+=1)l[o].d(1);l.length=n.length}},d(t){(0,s.RM)(l,t),t&&(0,s.og)(e)}}}function r(t){let e,n,l,a,o,i,r,d;function c(){t[6].call(n,t[15],t[16])}function u(){t[7].call(o,t[15],t[16])}return{c(){e=(0,s.bG)("div"),n=(0,s.bG)("input"),a=(0,s.Dh)(),o=(0,s.bG)("input"),i=(0,s.Dh)(),(0,s.Lj)(n,"type","text"),n.readOnly=l=!t[0].show_details,(0,s.Lj)(o,"type","number"),(0,s.Lj)(o,"class","number_input"),(0,s.Lj)(e,"class","item related"),(0,s.VH)(e,"show_details",t[0].show_details)},m(l,h){(0,s.$T)(l,e,h),(0,s.R3)(e,n),(0,s.Bm)(n,t[14].name),(0,s.R3)(e,a),(0,s.R3)(e,o),(0,s.Bm)(o,t[14].value),(0,s.R3)(e,i),r||(d=[(0,s.oL)(n,"input",c),(0,s.oL)(o,"input",u)],r=!0)},p(a,i){t=a,1&i&&l!==(l=!t[0].show_details)&&(n.readOnly=l),1&i&&n.value!==t[14].name&&(0,s.Bm)(n,t[14].name),1&i&&(0,s.sC)(o.value)!==t[14].value&&(0,s.Bm)(o,t[14].value),1&i&&(0,s.VH)(e,"show_details",t[0].show_details)},d(t){t&&(0,s.og)(e),r=!1,(0,s.j7)(d)}}}function d(t){let e,n,l;return{c(){e=(0,s.bG)("button"),e.textContent="add related keyword"},m(a,o){(0,s.$T)(a,e,o),n||(l=(0,s.oL)(e,"click",t[8]),n=!0)},p:s.ZT,d(t){t&&(0,s.og)(e),n=!1,l()}}}function c(t){let e,n,l,a,o;function i(){t[9].call(n,t[12],t[13])}return{c(){e=(0,s.bG)("div"),n=(0,s.bG)("input"),(0,s.Lj)(n,"type","text"),n.readOnly=l=!t[0].show_details,(0,s.Lj)(e,"class","item alias"),(0,s.VH)(e,"show_details",t[0].show_details)},m(l,r){(0,s.$T)(l,e,r),(0,s.R3)(e,n),(0,s.Bm)(n,t[11]),a||(o=(0,s.oL)(n,"input",i),a=!0)},p(a,o){t=a,1&o&&l!==(l=!t[0].show_details)&&(n.readOnly=l),1&o&&n.value!==t[11]&&(0,s.Bm)(n,t[11]),1&o&&(0,s.VH)(e,"show_details",t[0].show_details)},d(t){t&&(0,s.og)(e),a=!1,o()}}}function u(t){let e,n,l;return{c(){e=(0,s.bG)("button"),e.textContent="add alias"},m(a,o){(0,s.$T)(a,e,o),n||(l=(0,s.oL)(e,"click",t[10]),n=!0)},p:s.ZT,d(t){t&&(0,s.og)(e),n=!1,l()}}}function h(t){let e,n;return{c(){e=(0,s.bG)("h2"),n=(0,s.fL)(t[1]),(0,s.cz)(e,"color","red")},m(t,l){(0,s.$T)(t,e,l),(0,s.R3)(e,n)},p(t,e){2&e&&(0,s.rT)(n,t[1])},d(t){t&&(0,s.og)(e)}}}function m(t){let e,n,a,r,m,f,p,g,w,v,_,b,$,y,R,L,G,j,T,k,x,C,D=t[0].show_details?"save":"show/edit details";function N(t,e){return t[0].show_details?i:o}let O=N(t),B=O(t),S=t[0].show_details&&d(t),H=t[0].aliases,V=[];for(let e=0;e<H.length;e+=1)V[e]=c(l(t,H,e));let U=t[0].show_details&&u(t),Z=t[1]&&h(t);return{c(){e=(0,s.bG)("div"),n=(0,s.bG)("div"),a=(0,s.bG)("input"),r=(0,s.Dh)(),m=(0,s.bG)("div"),f=(0,s.bG)("h1"),f.textContent="related",p=(0,s.Dh)(),g=(0,s.bG)("div"),B.c(),w=(0,s.Dh)(),S&&S.c(),v=(0,s.Dh)(),_=(0,s.bG)("div"),b=(0,s.bG)("h1"),b.textContent="aliases",$=(0,s.Dh)(),y=(0,s.bG)("div");for(let t=0;t<V.length;t+=1)V[t].c();R=(0,s.Dh)(),U&&U.c(),L=(0,s.Dh)(),G=(0,s.bG)("button"),j=(0,s.fL)(D),T=(0,s.Dh)(),Z&&Z.c(),k=(0,s.cS)(),(0,s.Lj)(a,"type","text"),(0,s.Lj)(n,"class","name"),(0,s.Lj)(f,"class","mini-title"),(0,s.Lj)(g,"class","content"),(0,s.VH)(g,"show_details",t[0].show_details),(0,s.Lj)(m,"class","section related-section"),(0,s.Lj)(b,"class","mini-title"),(0,s.Lj)(y,"class","content"),(0,s.VH)(y,"show_details",t[0].show_details),(0,s.Lj)(_,"class","section alias-section"),(0,s.Lj)(e,"class","keyword-main")},m(l,o){(0,s.$T)(l,e,o),(0,s.R3)(e,n),(0,s.R3)(n,a),(0,s.Bm)(a,t[0].name),(0,s.R3)(e,r),(0,s.R3)(e,m),(0,s.R3)(m,f),(0,s.R3)(m,p),(0,s.R3)(m,g),B.m(g,null),(0,s.R3)(g,w),S&&S.m(g,null),(0,s.R3)(e,v),(0,s.R3)(e,_),(0,s.R3)(_,b),(0,s.R3)(_,$),(0,s.R3)(_,y);for(let t=0;t<V.length;t+=1)V[t].m(y,null);(0,s.R3)(y,R),U&&U.m(y,null),(0,s.R3)(e,L),(0,s.R3)(e,G),(0,s.R3)(G,j),(0,s.$T)(l,T,o),Z&&Z.m(l,o),(0,s.$T)(l,k,o),x||(C=[(0,s.oL)(a,"input",t[5]),(0,s.oL)(G,"click",t[3])],x=!0)},p(t,[e]){if(1&e&&a.value!==t[0].name&&(0,s.Bm)(a,t[0].name),O===(O=N(t))&&B?B.p(t,e):(B.d(1),B=O(t),B&&(B.c(),B.m(g,w))),t[0].show_details?S?S.p(t,e):(S=d(t),S.c(),S.m(g,null)):S&&(S.d(1),S=null),1&e&&(0,s.VH)(g,"show_details",t[0].show_details),1&e){let n;for(H=t[0].aliases,n=0;n<H.length;n+=1){const s=l(t,H,n);V[n]?V[n].p(s,e):(V[n]=c(s),V[n].c(),V[n].m(y,R))}for(;n<V.length;n+=1)V[n].d(1);V.length=H.length}t[0].show_details?U?U.p(t,e):(U=u(t),U.c(),U.m(y,null)):U&&(U.d(1),U=null),1&e&&(0,s.VH)(y,"show_details",t[0].show_details),1&e&&D!==(D=t[0].show_details?"save":"show/edit details")&&(0,s.rT)(j,D),t[1]?Z?Z.p(t,e):(Z=h(t),Z.c(),Z.m(k.parentNode,k)):Z&&(Z.d(1),Z=null)},i:s.ZT,o:s.ZT,d(t){t&&(0,s.og)(e),B.d(),S&&S.d(),(0,s.RM)(V,t),U&&U.d(),t&&(0,s.og)(T),Z&&Z.d(t),t&&(0,s.og)(k),x=!1,(0,s.j7)(C)}}}function f(t,e,n){let l,{data:a}=e,{all_keywords:o}=e,i="";return t.$$set=t=>{"data"in t&&n(0,a=t.data),"all_keywords"in t&&n(4,o=t.all_keywords)},t.$$.update=()=>{1&t.$$.dirty&&n(2,l=a.related.map(((t,e)=>`${t.name}:${t.value}`)).join(", "))},[a,i,l,()=>{if(a.show_details){const t=a.valid(o);t.valid?(n(0,a.show_details=!1,a),n(1,i="")):n(1,i=t.message)}else n(0,a.show_details=!0,a),n(1,i="")},o,function(){a.name=this.value,n(0,a)},function(t,e){t[e].name=this.value,n(0,a)},function(t,e){t[e].value=(0,s.sC)(this.value),n(0,a)},()=>{n(0,a.related=[...a.related,{name:"",value:null}],a)},function(t,e){t[e]=this.value,n(0,a)},()=>{n(0,a.aliases=[...a.aliases,""],a)}]}class p extends s.f_{constructor(t){super(),(0,s.S1)(this,t,f,m,s.N8,{data:0,all_keywords:4})}}const g=p;class w{constructor(t,e,n,s,l){this.name=t,this.aliases=e,this.related=n,this.category=s,this.show_details=l}static NewInst(){return new w("",[],[],!1,!0)}toJSON(){return JSON.stringify({name:this.name,aliases:this.aliases,related:this.related,category:this.category})}valid(t){if(!this.name)return{valid:!1,message:"missing keyword name"};if(t.filter((t=>t.name===this.name)).length>1)return{valid:!1,message:`keyword ${this.name} already exists`};let e=[];return this.related.forEach((n=>{let s=!1;t.forEach((t=>{t.name==this.name&&(s=!0)})),s||e.push(n.name)})),{valid:0===e.length,message:`invalid params: ${JSON.stringify(e)}`}}}function v(t,e,n){const s=t.slice();return s[6]=e[n],s}function _(t){let e,n;return e=new g({props:{data:t[6],all_keywords:t[0]}}),{c(){(0,s.YC)(e.$$.fragment)},m(t,l){(0,s.ye)(e,t,l),n=!0},p(t,n){const s={};1&n&&(s.data=t[6]),1&n&&(s.all_keywords=t[0]),e.$set(s)},i(t){n||((0,s.Ui)(e.$$.fragment,t),n=!0)},o(t){(0,s.et)(e.$$.fragment,t),n=!1},d(t){(0,s.vp)(e,t)}}}function b(t){let e,n,l,a,o,i,r,d,c,u,h,m,f,p=Object.values(t[0]).sort($),g=[];for(let e=0;e<p.length;e+=1)g[e]=_(v(t,p,e));const w=t=>(0,s.et)(g[t],1,1,(()=>{g[t]=null}));return{c(){e=(0,s.bG)("h1"),e.textContent="keyword mappings create/edit",n=(0,s.Dh)(),l=(0,s.bG)("div"),a=(0,s.bG)("button"),a.textContent="import and append mappings json",o=(0,s.Dh)(),i=(0,s.bG)("button"),i.textContent="export mappings json",r=(0,s.Dh)(),d=(0,s.bG)("div");for(let t=0;t<g.length;t+=1)g[t].c();c=(0,s.Dh)(),u=(0,s.bG)("button"),u.textContent="add keyword input",(0,s.Lj)(l,"class","top_buttons svelte-1oq552y"),(0,s.Lj)(d,"class","keywords svelte-1oq552y")},m(p,w){(0,s.$T)(p,e,w),(0,s.$T)(p,n,w),(0,s.$T)(p,l,w),(0,s.R3)(l,a),(0,s.R3)(l,o),(0,s.R3)(l,i),(0,s.$T)(p,r,w),(0,s.$T)(p,d,w);for(let t=0;t<g.length;t+=1)g[t].m(d,null);(0,s.R3)(d,c),(0,s.R3)(d,u),h=!0,m||(f=[(0,s.oL)(a,"click",t[1]),(0,s.oL)(u,"click",t[2])],m=!0)},p(t,[e]){if(1&e){let n;for(p=Object.values(t[0]).sort($),n=0;n<p.length;n+=1){const l=v(t,p,n);g[n]?(g[n].p(l,e),(0,s.Ui)(g[n],1)):(g[n]=_(l),g[n].c(),(0,s.Ui)(g[n],1),g[n].m(d,c))}for((0,s.dv)(),n=p.length;n<g.length;n+=1)w(n);(0,s.gb)()}},i(t){if(!h){for(let t=0;t<p.length;t+=1)(0,s.Ui)(g[t]);h=!0}},o(t){g=g.filter(Boolean);for(let t=0;t<g.length;t+=1)(0,s.et)(g[t]);h=!1},d(t){t&&(0,s.og)(e),t&&(0,s.og)(n),t&&(0,s.og)(l),t&&(0,s.og)(r),t&&(0,s.og)(d),(0,s.RM)(g,t),m=!1,(0,s.j7)(f)}}}const $=(t,e)=>t.name<e.name||!e.name;function y(t,e,n){let s=[];return[s,async()=>{const t=await undefined.files[fileInput.files.length-1].text();let e=JSON.parse(t),n=new w(e.name,e.aliases,e.related,e.category);s.push(n)},()=>{s.push(w.NewInst()),n(0,s=[...s])}]}class R extends s.f_{constructor(t){super(),(0,s.S1)(this,t,y,b,s.N8,{})}}const L=R;function G(t){let e,n,l,a;return l=new L({}),{c(){e=(0,s.bG)("title"),e.textContent="keyword input form",n=(0,s.Dh)(),(0,s.YC)(l.$$.fragment)},m(t,o){(0,s.$T)(t,e,o),(0,s.$T)(t,n,o),(0,s.ye)(l,t,o),a=!0},p:s.ZT,i(t){a||((0,s.Ui)(l.$$.fragment,t),a=!0)},o(t){(0,s.et)(l.$$.fragment,t),a=!1},d(t){t&&(0,s.og)(e),t&&(0,s.og)(n),(0,s.vp)(l,t)}}}class j extends s.f_{constructor(t){super(),(0,s.S1)(this,t,null,G,s.N8,{})}}const T=j}}]);