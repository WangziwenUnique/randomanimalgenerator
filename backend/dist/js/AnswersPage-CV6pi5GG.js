import{_ as e,a as n}from"./index-BmbeCsRI.js";import{a as s}from"./connections-answers-OLOdMENA.js";import{d as a,i as t,r as o,l as i,c as r,b as l,F as c,p as d,m as u,o as m,t as h,n as v}from"./vendor-MPsD2_ix.js";const p={class:"answers-page"},w={class:"answers-list"},g={class:"answer-header"},b={class:"answer-date"},y={class:"answer-desc"},C=["onClick"],f={key:0,class:"pagination"},k=["disabled"],T=["onClick"],N=["disabled"],Y=n(a({__name:"AnswersPage",setup(n){t((()=>{(()=>{const e=(e,n)=>{let s=document.querySelector(`meta[name="${e}"]`);s||(s=document.createElement("meta"),s.setAttribute("name",e),document.head.appendChild(s)),s.setAttribute("content",n)};document.title="Connections NYT Answers",e("description","Get daily answers and solutions for the NYT Connections Game. Find explanations and tips for solving word connection puzzles from The New York Times."),e("keywords","NYT Connections, Connections Game, Word Puzzle, Daily Answers, NYT Game Solutions, Word Connections")})()}));const a=o(1),Y=i((()=>{const e=Object.entries(s);return Object.fromEntries(e.sort((([e],[n])=>new Date(n).getTime()-new Date(e).getTime())))})),j=i((()=>{const e=Object.entries(Y.value),n=10*(a.value-1),s=n+10;return Object.fromEntries(e.slice(n,s))})),A=i((()=>Math.ceil(Object.keys(Y.value).length/10))),G=i((()=>{const e=[];let n=Math.max(1,a.value-2),s=Math.min(A.value,n+5-1);s-n+1<5&&(n=Math.max(1,s-5+1));for(let a=n;a<=s;a++)e.push(a);return e}));return(n,s)=>(m(),r("div",p,[s[3]||(s[3]=l("h1",{class:"page-title"},"Connections Answers",-1)),s[4]||(s[4]=l("div",{class:"game-intro"},[l("p",null,"NYT Connections is an innovative game that is reminiscent of Wordle and was created by The New York Times. In this game, players are challenged to discover groups of words connected by a common theme. The straightforward nature of this Word Connection Game is particularly appealing to us."),l("p",null,"Every day, we engage in playing and solving the NYT Connections, and we make sure to share all the answers, solutions, and hints, including today's answer, on this platform. We encourage you to play the daily NYT Connections game on New York Times website."),l("h2",{class:"section-title"},"Here are the most recent answers for the NYT Connections Game:")],-1)),l("div",w,[(m(!0),r(c,null,d(j.value,((a,t)=>(m(),r("div",{key:t,class:"answer-card"},[l("div",g,[l("h2",b,"Connections Answers for "+h(a.date),1),l("p",y," Looking for answers to the Connections on "+h(a.date)+"? We have a comprehensive answer guide ready to assist you in solving and unlocking the connections. ",1),l("button",{class:"read-more",onClick:e=>n.$router.push(`/connections-nyt-answers/${t}`)}," Read more ",8,C)]),s[2]||(s[2]=l("img",{src:e,alt:"Connections Game Logo",class:"answer-image"},null,-1))])))),128))]),A.value>1?(m(),r("div",f,[l("button",{class:"page-btn",disabled:1===a.value,onClick:s[0]||(s[0]=e=>a.value--)}," Previous ",8,k),(m(!0),r(c,null,d(G.value,(e=>(m(),r("button",{key:e,class:v(["page-btn",{active:a.value===e}]),onClick:n=>a.value=e},h(e),11,T)))),128)),l("button",{class:"page-btn",disabled:a.value===A.value,onClick:s[1]||(s[1]=e=>a.value++)}," Next ",8,N)])):u("",!0)]))}}),[["__scopeId","data-v-dc584bc1"]]);export{Y as default};
