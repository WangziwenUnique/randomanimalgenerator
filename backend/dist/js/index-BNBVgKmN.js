const __vite__mapDeps=(i,m=__vite__mapDeps,d=(m.f||(m.f=["js/AnswersPage-CD55egPs.js","js/connections-answers-DrY-xhui.js","js/vendor-CZ5gslDo.js","css/vendor-D-rRdhIO.css","css/AnswersPage-BRJ5XG72.css","js/AnswerDetailPage-DZdaIXNk.js","css/AnswerDetailPage-BWqAvz9h.css"])))=>i.map(i=>d[i]);
import{d as e,r as t,a,o as s,c as n,b as o,e as l,w as i,n as r,f as c,g as d,v as u,h as v,i as m,j as h,k as p,F as b,l as g,m as w,t as y,p as f,q as k,s as C,u as S,x as _,y as A,z as T,A as E}from"./vendor-CZ5gslDo.js";!function(){const e=document.createElement("link").relList;if(!(e&&e.supports&&e.supports("modulepreload"))){for(const e of document.querySelectorAll('link[rel="modulepreload"]'))t(e);new MutationObserver((e=>{for(const a of e)if("childList"===a.type)for(const e of a.addedNodes)"LINK"===e.tagName&&"modulepreload"===e.rel&&t(e)})).observe(document,{childList:!0,subtree:!0})}function t(e){if(e.ep)return;e.ep=!0;const t=function(e){const t={};return e.integrity&&(t.integrity=e.integrity),e.referrerPolicy&&(t.referrerPolicy=e.referrerPolicy),"use-credentials"===e.crossOrigin?t.credentials="include":"anonymous"===e.crossOrigin?t.credentials="omit":t.credentials="same-origin",t}(e);fetch(e.href,t)}}();const q="/logo.svg",G={class:"mobile-header"},N={class:"mobile-header-top"},O={class:"mobile-nav"},P={class:"mobile-search"},x=(e,t)=>{const a=e.__vccOpts||e;for(const[s,n]of t)a[s]=n;return a},L=x(e({__name:"MobileHeader",setup(e){const m=t(!1),h=t(""),p=()=>{m.value=!m.value},b=()=>{m.value=!1},g=()=>{b()};return(e,t)=>{const w=a("router-link");return s(),n("header",G,[o("div",N,[l(w,{to:"/",class:"mobile-logo"},{default:i((()=>t[1]||(t[1]=[o("img",{src:q,alt:"Connections Game",class:"mobile-logo-image"},null,-1)]))),_:1}),o("button",{class:"menu-button",onClick:p},[o("span",{class:r(["menu-icon",{"is-active":m.value}])},null,2)])]),o("div",{class:r(["mobile-menu",{"is-open":m.value}])},[o("nav",O,[l(w,{to:"/daily",class:r(["mobile-nav-item",{active:"/daily"===e.$route.path}]),onClick:b},{default:i((()=>t[2]||(t[2]=[c(" Connections ")]))),_:1},8,["class"]),l(w,{to:"/connections-nyt-answers",class:r(["mobile-nav-item",{active:e.$route.path.includes("connections-nyt-answers")}]),onClick:b},{default:i((()=>t[3]||(t[3]=[c(" Connections NYT Answers ")]))),_:1},8,["class"])]),o("div",P,[d(o("input",{type:"text","onUpdate:modelValue":t[0]||(t[0]=e=>h.value=e),placeholder:"Search...",onKeyup:v(g,["enter"])},null,544),[[u,h.value]]),o("button",{class:"mobile-search-button",onClick:g},t[4]||(t[4]=[o("svg",{xmlns:"http://www.w3.org/2000/svg",width:"16",height:"16",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor","stroke-width":"2","stroke-linecap":"round","stroke-linejoin":"round"},[o("circle",{cx:"11",cy:"11",r:"8"}),o("line",{x1:"21",y1:"21",x2:"16.65",y2:"16.65"})],-1)]))])],2)])}}}),[["__scopeId","data-v-2cf26cc4"]]),I={key:1,class:"header"},M={class:"header-container"},D={class:"logo"},j={class:"nav-menu"},R={class:"right-section"},Y={class:"search-box"},$=x(e({__name:"TheHeader",setup(e){const b=t(!1),g=()=>{b.value=window.innerWidth<500};m((()=>{g(),window.addEventListener("resize",g)})),h((()=>{window.removeEventListener("resize",g)}));const w=t(""),y=()=>{};return(e,t)=>{const m=a("router-link");return b.value?(s(),p(L,{key:0})):(s(),n("header",I,[o("div",M,[o("div",D,[l(m,{to:"/"},{default:i((()=>t[1]||(t[1]=[o("img",{src:q,alt:"Connections Game",class:"logo-image"},null,-1)]))),_:1})]),o("nav",j,[l(m,{to:"/daily",class:r(["nav-item",{active:"/daily"===e.$route.path}])},{default:i((()=>t[2]||(t[2]=[c(" Connections ")]))),_:1},8,["class"])]),o("div",R,[l(m,{to:"/connections-nyt-answers",class:r(["answers-link",{active:e.$route.path.includes("connections-nyt-answers")}])},{default:i((()=>t[3]||(t[3]=[c(" Connections NYT Answers ")]))),_:1},8,["class"]),o("div",Y,[d(o("input",{type:"text","onUpdate:modelValue":t[0]||(t[0]=e=>w.value=e),placeholder:"Search...",onKeyup:v(y,["enter"])},null,544),[[u,w.value]]),o("button",{class:"search-button",onClick:y},t[4]||(t[4]=[o("svg",{xmlns:"http://www.w3.org/2000/svg",width:"16",height:"16",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor","stroke-width":"2","stroke-linecap":"round","stroke-linejoin":"round"},[o("circle",{cx:"11",cy:"11",r:"8"}),o("line",{x1:"21",y1:"21",x2:"16.65",y2:"16.65"})],-1)]))])])])]))}}}),[["__scopeId","data-v-cc462a2e"]]),H={class:"main-content"},z=e({__name:"App",setup:e=>(e,t)=>{const i=a("router-view");return s(),n(b,null,[l($),o("main",H,[l(i)])],64)}}),F={},W=function(e,t,a){let s=Promise.resolve();if(t&&t.length>0){document.getElementsByTagName("link");const e=document.querySelector("meta[property=csp-nonce]"),a=(null==e?void 0:e.nonce)||(null==e?void 0:e.getAttribute("nonce"));s=Promise.allSettled(t.map((e=>{if((e=function(e){return"/"+e}(e))in F)return;F[e]=!0;const t=e.endsWith(".css"),s=t?'[rel="stylesheet"]':"";if(document.querySelector(`link[href="${e}"]${s}`))return;const n=document.createElement("link");return n.rel=t?"stylesheet":"modulepreload",t||(n.as="script"),n.crossOrigin="",n.href=e,a&&n.setAttribute("nonce",a),document.head.appendChild(n),t?new Promise(((t,a)=>{n.addEventListener("load",t),n.addEventListener("error",(()=>a(new Error(`Unable to preload CSS for ${e}`))))})):void 0})))}function n(e){const t=new Event("vite:preloadError",{cancelable:!0});if(t.payload=e,window.dispatchEvent(t),!t.defaultPrevented)throw e}return s.then((t=>{for(const e of t||[])"rejected"===e.status&&n(e.reason);return e().catch(n)}))},K="/images/look_for_a_common_thing.jpeg",U={class:"connections-game"},B={key:0,class:"start-game-container"},V={class:"timer-container"},J={class:"game-area"},Q={class:"found-groups"},X={class:"group-title"},Z={class:"group-words"},ee={key:0,class:"loading-container"},te={key:1,class:"empty-state"},ae={key:2,class:"game-grid"},se=["data-length","onClick","disabled"],ne={key:3,class:"game-over"},oe={key:4,class:"game-controls"},le=["disabled"],ie={key:5,class:"game-status"},re=[{path:"/",redirect:"/daily"},{path:"/daily",name:"Daily",component:x(e({__name:"ConnectionsGame",setup(e){const a=t([{title:"MUSICAL NOTES",words:["FA","MI","LA","SOL"],color:"yellow"},{title:"EMOTIONS",words:["FEAR","ANGER","SADNESS","JOY"],color:"green"},{title:"SHADES",words:["GREY","CHARCOAL","TAN","TAPESTRY"],color:"blue"},{title:"SPORTS EQUIPMENT",words:["RACK","SKEET","PADDLE","NET"],color:"purple"}]),l=t([]),i=t([]),c=t("Select four related words"),d=t(!1),u=t(4),v=t([]),p="https://connectionshinttoday.today",C=t(!0),S=t(""),_=t(0);let A=null;const T=()=>{A||(A=window.setInterval((()=>{F.value||_.value++}),1e3))},E=()=>{A&&(clearInterval(A),A=null)},q=()=>{_.value=0,E()},G=async()=>{C.value=!0;try{const e=await fetch(`${p}/api/connections/words`),t=await e.json();t.words&&Array.isArray(t.words)&&t.words.length>0?(v.value=t.words.sort((()=>Math.random()-.5)),t.gameId&&(S.value=t.gameId)):v.value=[]}catch(e){v.value=[]}finally{C.value=!1}},N=g((()=>{const e=new Set(l.value.flatMap((e=>e.words)));return a.value.flatMap((e=>e.words)).filter((t=>!e.has(t)))})),O=t(!1),P=async()=>{O.value=!0,q(),v.value.length>0?T():await x()},x=async()=>{await G(),v.value.length>0&&T()},L=e=>l.value.some((t=>t.words.includes(e))),I=t(!1);t(null);const M=async()=>{if(4===i.value.length&&!I.value)try{I.value=!0;const e=await fetch(`${p}/api/connections/submit`,{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({gameID:S.value,words:i.value})}),t=await e.json();if(t.success){const e={title:t.category,words:i.value,color:j()};v.value=v.value.filter((e=>!i.value.includes(e))),l.value.push(e),i.value=[],4===l.value.length?c.value="Congratulations! You completed the game!":c.value="Select four related words",d.value=!1}else u.value>0&&(u.value--,c.value="These words don't form a group",d.value=!0,0===u.value&&(c.value="Game Over"))}catch(e){u.value>0&&(u.value--,c.value="Failed to submit. Please try again.",d.value=!0,0===u.value&&(c.value="Game Over"))}finally{setTimeout((()=>{I.value=!1}),500)}},D=["yellow","green","blue","purple"],j=()=>D[l.value.length]||D[0],R=()=>{i.value=[]},Y=async()=>{q(),i.value=[],u.value=4,c.value="Select four related words",d.value=!1,l.value=[],await G(),v.value.length>0&&T()},$=g((()=>0===u.value)),H=()=>{q(),v.value=[...N.value].sort((()=>Math.random()-.5)),i.value=[],u.value=4,c.value="Select four related words",d.value=!1,l.value=[],T()},z=async()=>{q(),await x()};m((()=>{}));const F=g((()=>{const e=4===l.value.length;return e&&E(),e}));return h((()=>{E()})),(e,t)=>{return s(),n("div",U,[t[9]||(t[9]=o("h1",{class:"game-title"},"Connections Hint Today",-1)),t[10]||(t[10]=o("p",{class:"game-subtitle"},"Daily NYT Connections Game Help & Hints",-1)),O.value?w("",!0):(s(),n("div",B,[o("div",{class:"start-game-content"},[t[0]||(t[0]=o("h2",{class:"start-game-title"},"Ready to Play?",-1)),t[1]||(t[1]=o("h3",{class:"start-game-subtitle"},"Daily Word Connection Challenge",-1)),t[2]||(t[2]=o("p",{class:"start-game-desc"},"Find the connections between words and form perfect groups!",-1)),o("button",{class:"play-button",onClick:P}," Play Now ")])])),O.value?(s(),n(b,{key:1},[o("div",V,[o("div",{class:r(["timer",{"timer-complete":F.value}])}," Time: "+y((a=_.value,`${Math.floor(a/60)}:${(a%60).toString().padStart(2,"0")}`)),3)]),o("div",J,[o("div",Q,[(s(!0),n(b,null,f(l.value,((e,t)=>(s(),n("div",{key:"found-"+t,class:r(["found-group",e.color])},[o("h3",X,y(e.title),1),o("p",Z,y(e.words.join(", ")),1)],2)))),128))]),C.value?(s(),n("div",ee,t[3]||(t[3]=[o("div",{class:"loading-spinner"},null,-1),o("p",{class:"loading-text"},"Loading words...",-1)]))):v.value.length||F.value?(s(),n("div",ae,[(s(!0),n(b,null,f(v.value,((e,t)=>(s(),n("button",{key:t,class:r(["word-tile",{selected:i.value.includes(e)}]),"data-length":e.length>8?"long":"normal",onClick:t=>(e=>{const t=i.value.indexOf(e);-1===t?i.value.length<4&&i.value.push(e):i.value.splice(t,1)})(e),disabled:$.value||L(e)},y(e),11,se)))),128))])):(s(),n("div",te,[o("div",{class:"empty-state-content"},[t[4]||(t[4]=o("svg",{class:"empty-state-icon",viewBox:"0 0 24 24",fill:"none",stroke:"currentColor","stroke-width":"2"},[o("path",{d:"M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"})],-1)),t[5]||(t[5]=o("h3",{class:"empty-state-title"},"No Words Available",-1)),t[6]||(t[6]=o("p",{class:"empty-state-message"},"Unable to load words. Please try again later.",-1)),o("button",{class:"retry-button",onClick:z}," Try Again ")])])),$.value?(s(),n("div",ne,[t[7]||(t[7]=o("h2",{class:"game-over-title"},"Game Over!",-1)),t[8]||(t[8]=o("h3",{class:"game-over-subtitle"},"All attempts used up!",-1)),o("div",{class:"game-over-buttons"},[o("button",{class:"game-over-btn",onClick:H}," Restart the same Game "),o("button",{class:"game-over-btn",onClick:Y}," Start New Game ")])])):(s(),n("div",oe,[o("button",{class:"submit-btn",onClick:M,disabled:4!==i.value.length||I.value},y(I.value?"Submitting...":"Submit"),9,le),o("button",{class:"deselect-btn",onClick:R}," Deselect "),o("button",{class:"new-game-btn",onClick:Y}," Start New Game ")])),$.value?w("",!0):(s(),n("div",ie,[o("p",{class:r({error:d.value})},y(c.value),3),o("p",null,"Mistakes remaining: "+y(u.value),1)]))])],64)):w("",!0),t[11]||(t[11]=k('<div class="game-introduction-wrapper" data-v-c29ebb17><div class="game-introduction" data-v-c29ebb17><h2 class="intro-title" data-v-c29ebb17>NYT Connections Game Guide</h2><h3 class="intro-subtitle" data-v-c29ebb17>Today&#39;s Strategy &amp; Tips</h3><p class="intro-text" data-v-c29ebb17> Welcome to Connections Hint Today, your daily companion for the NYT Connections Game. Get help finding common threads between words and form four perfect groups while limiting mistakes to four. Our daily updated hints and strategic guidance will help you master the Connections puzzle. </p><h2 class="how-to-play-title" data-v-c29ebb17>Game Instructions</h2><h3 class="how-to-play-subtitle" data-v-c29ebb17>Step by Step Guide</h3><div class="how-to-play-steps" data-v-c29ebb17><div class="step" data-v-c29ebb17><div class="step-content" data-v-c29ebb17><h3 class="step-title" data-v-c29ebb17>1. Read the words</h3><p data-v-c29ebb17>The first step is to carefully read and understand the words provided. Take your time to understand each word and think about what it means in relation to the Connections Puzzle.</p></div><img src="/images/read_the_words.jpeg" alt="Read the words" class="step-image" data-v-c29ebb17></div><div class="step" data-v-c29ebb17><div class="step-content" data-v-c29ebb17><h3 class="step-title" data-v-c29ebb17>2. Look for a Common Thing</h3><p data-v-c29ebb17>Find the common theme that connects them. Ask yourself: Do the words belong to the same group? Are they similar in some way? Do they have a connection to a specific topic or idea?</p></div><img src="'+K+'" alt="Look for connections" class="step-image" data-v-c29ebb17></div><div class="step" data-v-c29ebb17><div class="step-content" data-v-c29ebb17><h3 class="step-title" data-v-c29ebb17>3. Select &amp; Submit Your Answer</h3><p data-v-c29ebb17>Once you have identified the common theme and found the four words that fit within that theme, select them and submit your answer using the Submit button.</p></div><img src="/images/select_submit_your_answer.jpeg" alt="Submit your answer" class="step-image" data-v-c29ebb17></div></div><h2 class="faq-section-title" data-v-c29ebb17>Help &amp; Support</h2><h3 class="faq-title" data-v-c29ebb17>Frequently Asked Questions</h3><div class="faq-list" data-v-c29ebb17><div class="faq-item" data-v-c29ebb17><h3 class="faq-question" data-v-c29ebb17>What is Connections Game?</h3><p data-v-c29ebb17>Connections is a daily game where players find common threads between words. The goal is to select four groups of four words, while being mindful of making no more than four mistakes.</p></div><div class="faq-item" data-v-c29ebb17><h3 class="faq-question" data-v-c29ebb17>What are the rules of the Connections Game?</h3><ul class="faq-rules" data-v-c29ebb17><li data-v-c29ebb17>Sort 16 Words into 4 Groups</li><li data-v-c29ebb17>Find connections between the words to determine which ones belong together</li><li data-v-c29ebb17>Select four words at a time and submit your answer</li><li data-v-c29ebb17>If your selection is correct, the words will be placed in their respective group</li><li data-v-c29ebb17>Be cautious, as making four mistakes will result in losing the game</li><li data-v-c29ebb17>You can only play once per day</li></ul></div><div class="faq-item" data-v-c29ebb17><h3 class="faq-question" data-v-c29ebb17>How Many Mistakes Can I Make?</h3><p data-v-c29ebb17>You are allowed to make up to four mistakes in the Connections before the round ends.</p></div><div class="faq-item" data-v-c29ebb17><h3 class="faq-question" data-v-c29ebb17>Can Kids Play Connections Game?</h3><p data-v-c29ebb17>Yes, Kids can play Connections on their Mobile and Laptop. It&#39;s a fun word game that can help them learn and think logically. The game has different difficulty levels, so they can start with easier categories and work their way up.</p></div><div class="faq-item" data-v-c29ebb17><h3 class="faq-question" data-v-c29ebb17>Are hints available?</h3><p data-v-c29ebb17>No, we don&#39;t have hints in the Connections game. If you are stuck and unable to find answers, you can use the &quot;Reveal Answers&quot; button. This will reveal all the answers.</p></div></div></div></div>',1))]);var a}}}),[["__scopeId","data-v-c29ebb17"]])},{path:"/connections-nyt-answers",name:"Answers",component:()=>W((()=>import("./AnswersPage-CD55egPs.js")),__vite__mapDeps([0,1,2,3,4]))},{path:"/connections-nyt-answers/:date",name:"AnswerDetail",component:()=>W((()=>import("./AnswerDetailPage-DZdaIXNk.js")),__vite__mapDeps([5,2,3,1,6]))},{path:"/sitemap.xml",name:"sitemap",component:{render(){}},beforeEnter:(e,t,a)=>(window.location.href="/sitemap.xml",!1)}],ce=C({history:S(),routes:re});ce.beforeEach(((e,t,a)=>{e.meta.title&&(document.title=e.meta.title),a()}));const de=_({components:T,directives:E,theme:{defaultTheme:"light"}}),ue=A(z);ue.use(de),ue.use(ce),ue.mount("#app");export{K as _,x as a};
