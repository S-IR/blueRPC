type bluerpc ={depth1:{depth2:{test:{query:(queryParams:{ Something: string,})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});mutation:(input : {queryParams:{ Something: string,},input:{ house: string,}})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});}},zap:{mutation:(input : {queryParams:{ Something: string,},input:{ house: string,}})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});}},error:{mutation:(input : {queryParams:{ Something: string,},input:{ house: string,}})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});},hello:{world:{query:()=>({ Something: string|undefined,});}},helloWorld:{query:(queryParams:{ Something: string,})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});},test:{query:(queryParams:{ Something: string,})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});mutation:(input : {queryParams:{ Something: string,},input:{ house: string,}})=>({ fieldOneOut: string, fieldTwoOut: string|undefined, fieldThreeOut: string,});}}