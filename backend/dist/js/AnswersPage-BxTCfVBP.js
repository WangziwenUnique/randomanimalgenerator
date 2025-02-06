import{_ as e,a as s}from"./index-hEc_k9Yx.js";import{a}from"./connections-answers-DrY-xhui.js";import{d as n,r as t,l as o,c as i,b as r,F as l,p as c,m as d,o as u,t as m,n as v}from"./vendor-e11xv0F6.js";const g={class:"answers-page"},h={class:"answers-list"},p={class:"answer-header"},w={class:"answer-date"},b={class:"answer-desc"},y=["onClick"],f={key:0,class:"pagination"},k=["disabled"],C=["onClick"],T=["disabled"],j=s(n({__name:"AnswersPage",setup(s){const n=t(1),j=o((()=>{const e=Object.entries(a);return Object.fromEntries(e.sort((([e],[s])=>new Date(s).getTime()-new Date(e).getTime())))})),N=o((()=>{const e=Object.entries(j.value),s=10*(n.value-1),a=s+10;return Object.fromEntries(e.slice(s,a))})),Y=o((()=>Math.ceil(Object.keys(j.value).length/10))),x=o((()=>{const e=[];let s=Math.max(1,n.value-2),a=Math.min(Y.value,s+5-1);a-s+1<5&&(s=Math.max(1,a-5+1));for(let n=s;n<=a;n++)e.push(n);return e}));return(s,a)=>(u(),i("div",g,[a[3]||(a[3]=r("h1",{class:"page-title"},"Connections Answers",-1)),a[4]||(a[4]=r("div",{class:"game-intro"},[r("p",null,"NYT Connections is an innovative game that is reminiscent of Wordle and was created by The New York Times. In this game, players are challenged to discover groups of words connected by a common theme. The straightforward nature of this Word Connection Game is particularly appealing to us."),r("p",null,"Every day, we engage in playing and solving the NYT Connections, and we make sure to share all the answers, solutions, and hints, including today's answer, on this platform. We encourage you to play the daily NYT Connections game on New York Times website."),r("h2",{class:"section-title"},"Here are the most recent answers for the NYT Connections Game:")],-1)),r("div",h,[(u(!0),i(l,null,c(N.value,((n,t)=>(u(),i("div",{key:t,class:"answer-card"},[r("div",p,[r("h2",w,"Connections Answers for "+m(n.date),1),r("p",b," Looking for answers to the Connections on "+m(n.date)+"? We have a comprehensive answer guide ready to assist you in solving and unlocking the connections. ",1),r("button",{class:"read-more",onClick:e=>s.$router.push(`/connections-nyt-answers/${t}`)}," Read more ",8,y)]),a[2]||(a[2]=r("img",{src:e,alt:"Connections Game Logo",class:"answer-image"},null,-1))])))),128))]),Y.value>1?(u(),i("div",f,[r("button",{class:"page-btn",disabled:1===n.value,onClick:a[0]||(a[0]=e=>n.value--)}," Previous ",8,k),(u(!0),i(l,null,c(x.value,(e=>(u(),i("button",{key:e,class:v(["page-btn",{active:n.value===e}]),onClick:s=>n.value=e},m(e),11,C)))),128)),r("button",{class:"page-btn",disabled:n.value===Y.value,onClick:a[1]||(a[1]=e=>n.value++)}," Next ",8,T)])):d("",!0)]))}}),[["__scopeId","data-v-6e623ee9"]]);export{j as default};
