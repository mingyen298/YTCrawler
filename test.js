var PB={
Pn:function(a){a.reverse()},
Tg:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
Tn:function(a,b){a.splice(0,b)}};
Jwa=function(a){a=a.split("");PB.Tg(a,48);PB.Tg(a,43);PB.Tn(a,2);PB.Pn(a,79);PB.Tn(a,2);PB.Tg(a,39);PB.Tg(a,19);PB.Tn(a,1);PB.Pn(a,28);return a.join("")};


fake = "AOq0QJ8wRgIhAIOIP_b4Nnu_wAH7PVeDxOTxhME9xB5ldu763vDc6tO-AiEA6bHjKaCbBC7SegbwoRgu6ws_5BnX2h3VBMQYtLP6rg4="
real = Jwa(fake)

console.log(real)