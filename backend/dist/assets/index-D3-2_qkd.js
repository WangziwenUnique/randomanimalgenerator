import{bL as O,bM as z,f as C,bS as v,bT as S,bW as j,bX as w,bY as _,aj as x,i as L,k as R,bZ as T,b_ as b}from"./vendor-CdkM_rDc.js";import{_ as A,R as E,f as P,C as I,I as M,g as V,T as N,h as g}from"./image-resize-BwgWHxnz.js";import{C as k}from"./image-crop-fIVVII7d.js";(function(){const e=document.createElement("link").relList;if(e&&e.supports&&e.supports("modulepreload"))return;for(const n of document.querySelectorAll('link[rel="modulepreload"]'))o(n);new MutationObserver(n=>{for(const r of n)if(r.type==="childList")for(const c of r.addedNodes)c.tagName==="LINK"&&c.rel==="modulepreload"&&o(c)}).observe(document,{childList:!0,subtree:!0});function s(n){const r={};return n.integrity&&(r.integrity=n.integrity),n.referrerPolicy&&(r.referrerPolicy=n.referrerPolicy),n.crossOrigin==="use-credentials"?r.credentials="include":n.crossOrigin==="anonymous"?r.credentials="omit":r.credentials="same-origin",r}function o(n){if(n.ep)return;n.ep=!0;const r=s(n);fetch(n.href,r)}})();const q={class:"page-container"},B=O({__name:"App",setup(t){return(e,s)=>{const o=v("router-view");return S(),z("div",q,[C(o)])}}}),D=A(B,[["__scopeId","data-v-e12ef009"]]),W=[{path:"/",redirect:"/resize"},{path:"/resize",name:"resize",component:E},{path:"/crop",name:"crop",component:k}],F=j({history:w(),routes:W});function H(t,e){const s=e.modifiers||{},o=e.value,{once:n,immediate:r,...c}=s,a=!Object.keys(c).length,{handler:u,options:i}=typeof o=="object"?o:{handler:o,options:{attributes:(c==null?void 0:c.attr)??a,characterData:(c==null?void 0:c.char)??a,childList:(c==null?void 0:c.child)??a,subtree:(c==null?void 0:c.sub)??a}},l=new MutationObserver(function(){let f=arguments.length>0&&arguments[0]!==void 0?arguments[0]:[],p=arguments.length>1?arguments[1]:void 0;u==null||u(f,p),n&&h(t,e)});r&&(u==null||u([],l)),t._mutate=Object(t._mutate),t._mutate[e.instance.$.uid]={observer:l},l.observe(t,i)}function h(t,e){var s;(s=t._mutate)!=null&&s[e.instance.$.uid]&&(t._mutate[e.instance.$.uid].observer.disconnect(),delete t._mutate[e.instance.$.uid])}const X={mounted:H,unmounted:h};function Y(t,e){var n,r;const s=e.value,o={passive:!((n=e.modifiers)!=null&&n.active)};window.addEventListener("resize",s,o),t._onResize=Object(t._onResize),t._onResize[e.instance.$.uid]={handler:s,options:o},(r=e.modifiers)!=null&&r.quiet||s()}function Z(t,e){var n;if(!((n=t._onResize)!=null&&n[e.instance.$.uid]))return;const{handler:s,options:o}=t._onResize[e.instance.$.uid];window.removeEventListener("resize",s,o),delete t._onResize[e.instance.$.uid]}const G={mounted:Y,unmounted:Z};function y(t,e){const{self:s=!1}=e.modifiers??{},o=e.value,n=typeof o=="object"&&o.options||{passive:!0},r=typeof o=="function"||"handleEvent"in o?o:o.handler,c=s?t:e.arg?document.querySelector(e.arg):window;c&&(c.addEventListener("scroll",r,n),t._onScroll=Object(t._onScroll),t._onScroll[e.instance.$.uid]={handler:r,options:n,target:s?void 0:c})}function $(t,e){var r;if(!((r=t._onScroll)!=null&&r[e.instance.$.uid]))return;const{handler:s,options:o,target:n=t}=t._onScroll[e.instance.$.uid];n.removeEventListener("scroll",s,o),delete t._onScroll[e.instance.$.uid]}function J(t,e){e.value!==e.oldValue&&($(t,e),y(t,e))}const Q={mounted:y,unmounted:$,updated:J};function U(t,e){const s=typeof t=="string"?v(t):t,o=K(s,e);return{mounted:o,updated:o,unmounted(n){_(null,n)}}}function K(t,e){return function(s,o,n){var f,p,m;const r=typeof e=="function"?e(o):e,c=((f=o.value)==null?void 0:f.text)??o.value??(r==null?void 0:r.text),a=x(o.value)?o.value:{},u=()=>c??s.textContent,i=(n.ctx===o.instance.$?(p=ee(n,o.instance.$))==null?void 0:p.provides:(m=n.ctx)==null?void 0:m.provides)??o.instance.$.provides,l=L(t,R(r,a),u);l.appContext=Object.assign(Object.create(null),o.instance.$.appContext,{provides:i}),_(l,s)}}function ee(t,e){const s=new Set,o=r=>{var c,a;for(const u of r){if(!u)continue;if(u===t||u.el&&t.el&&u.el===t.el)return!0;s.add(u);let i;if(u.suspense?i=o([u.ssContent]):Array.isArray(u.children)?i=o(u.children):(c=u.component)!=null&&c.vnode&&(i=o([(a=u.component)==null?void 0:a.subTree])),i)return i;s.delete(u)}return!1};if(!o([e.subTree]))return e;const n=Array.from(s).reverse();for(const r of n)if(r.component)return r.component;return e}const te=U(P,t=>{var e;return{activator:"parent",location:(e=t.arg)==null?void 0:e.replace("-"," "),text:typeof t.value=="boolean"?void 0:t.value}}),oe=Object.freeze(Object.defineProperty({__proto__:null,ClickOutside:I,Intersect:M,Mutate:X,Resize:G,Ripple:V,Scroll:Q,Tooltip:te,Touch:N},Symbol.toStringTag,{value:"Module"})),ne=T({components:g,directives:oe,theme:{defaultTheme:"light"}}),d=b(D);d.use(ne);d.use(F);d.mount("#app");
