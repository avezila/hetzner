#!/bin/node
require('console.table');
let fs = require('fs');
let md = require('node-utf-format');
let zlib = require("zlib");
let request = require('sync-request');
const Telegram = require('telegraf/telegram')



const token = '787991996:AAEC0_MHueKV4bETpzZb3ZRfPa2wqfgATX0';

const bot = new Telegram(token)


let source = fs.readFileSync('/opt/hetzner.gz');
source = zlib.gunzipSync(source);
source = zlib.gunzipSync(source);
source += ']';

let all_dumps = JSON.parse(source).reduce((a, b) => [...a, ...b.server], []);
// let last_dump = JSON.parse(source).pop().server;
let last_dump = request('GET', 'https://www.hetzner.com/a_hz_serverboerse/live_data.json?m=1539681284257', {
  headers: {
    'user-agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36',
  },
}).getBody();
last_dump = JSON.parse(last_dump).server;


let cpub = 454*1.032;
let hddb = 141*1.08;
let ssdb = 14.1/0.979;
let ramb = 3.35*1.022;
function Key(s) {
	return `${s.ram}${s.cpu_benchmark}${s.hdd_size}${s.description.join(' ')}`;
}
function A (a,b,c) {
	return Math.sqrt(a*a+b*b+c*c)/3;
}
function calc(s) {
	let all = [s]
		.map(s => ({c: s.cpu_benchmark/s.price/cpub, ...s}))
		.map(s => ({h: s.hdd_size/s.price/hddb, ...s}))
		.map(s => ({r: s.ram/s.price/ramb, ...s}))
		.map(s => {
			if (!s.freetext.match(/hdd/i) && s.freetext.match(/ssd/i)) {s.h *= ssdb;}
			return s;
		})
		.map(s => ({a: A(s.h,s.r,s.c), ...s}))
		.map(s => ({
                        ...s,
                        cpub: Math.round(s.c*100*10)/10,
                        hddb: Math.round(s.h*100*10)/10,
                        ramb: Math.round(s.r*100*10)/10,
                        allb: Math.round(s.a*100*10)/10,
                }))
                .map(s => ({
                        allb: s.allb>=40 ? md.format(`${s.allb}##`) : s.allb,
                        cpub: s.cpub>=90 ? md.format(`${s.cpub}##`) : s.cpub,
                        hddb: s.hddb>=90 ? md.format(`${s.hddb}##`) : s.hddb,
                        ramb: s.ramb>=90 ? md.format(`${s.ramb}##`) : s.ramb,
			...s,
		}));
	return all[0];
}
function print (key, dumps) {
	console.log(`${key}:`)
	let map = new Map();
	dumps
		.map(s => ({
			...s,
			next_reduce: s.next_reduce === 0 ? Infinity : s.next_reduce,
		}))
		.sort((a,b) => a.next_reduce - b.next_reduce)
		.sort((a,b) => Number(a.price) - Number(b.price))
		.forEach(s => {
			if (map.has(Key(s))) return;
			map.set(Key(s), s);
		});

	let servers = Array.from(map.values());
	let hdd = servers
		.map(s => ({c: s.cpu_benchmark/s.price/cpub, ...s}))
		.map(s => ({h: s.hdd_size/s.price/hddb, ...s}))
		.map(s => ({r: s.ram/s.price/ramb, ...s}))
		.map(s => {
			if (!s.freetext.match(/hdd/i) && s.freetext.match(/ssd/i)) {s.h *= ssdb;}
			return s;
		})
		.map(s => ({a: A(s.h,s.r,s.c), ...s}))
		.sort((a,b) => b.h - a.h)
		.filter(s =>
			!s.freetext.match(/ssd/i) &&
			s.freetext.match(/hdd/i)
		)
		.slice(0, 20)

	let ssd = servers
		.map(s => ({c: s.cpu_benchmark/s.price/cpub, ...s}))
		.map(s => ({h: s.hdd_size/s.price/hddb, ...s}))
		.map(s => ({r: s.ram/s.price/ramb, ...s}))
		.map(s => {
			if (!s.freetext.match(/hdd/i) && s.freetext.match(/ssd/i)) {s.h *= ssdb;}
			return s;
		})
		.map(s => ({a: A(s.h,s.r,s.c), ...s}))
		.sort((a,b) => b.h - a.h)
		.filter(s =>
			s.freetext.match(/ssd/i) &&
			!s.freetext.match(/hdd/i)
		)
		.slice(0, 20)

	let cpu = servers
		.map(s => ({c: s.cpu_benchmark/s.price/cpub, ...s}))
		.map(s => ({r: s.ram/s.price/ramb, ...s}))
		.map(s => ({h: s.hdd_size/s.price/hddb, ...s}))
		.map(s => {
			if (!s.freetext.match(/hdd/i) && s.freetext.match(/ssd/i)) {s.h *= ssdb;}
			return s;
		})
		.map(s => ({a: A(s.h,s.r,s.c), ...s}))
		.sort((a,b) => b.c - a.c)
		.slice(0, 20)
	
	let ram = servers
		.map(s => ({c: s.cpu_benchmark/s.price/cpub, ...s}))
		.map(s => ({r: s.ram/s.price/ramb, ...s}))
		.map(s => ({h: s.hdd_size/s.price/hddb, ...s}))
		.map(s => {
			if (!s.freetext.match(/hdd/i) && s.freetext.match(/ssd/i)) {s.h *= ssdb;}
			return s;
		})
		.map(s => ({a: A(s.h,s.r,s.c), ...s}))
		.sort((a,b) => b.r - a.r)
		.slice(0, 20)
	
	let all = servers
		.map(s => ({c: s.cpu_benchmark/s.price/cpub, ...s}))
		.map(s => ({h: s.hdd_size/s.price/hddb, ...s}))
		.map(s => ({r: s.ram/s.price/ramb, ...s}))
		.map(s => {
			if (!s.freetext.match(/hdd/i) && s.freetext.match(/ssd/i)) {s.h *= ssdb;}
			return s;
		})
		.map(s => ({a: A(s.h,s.r,s.c), ...s}))
		.sort((a,b) => b.a - a.a)
		.slice(0, 20)

	pretty("ram", ram);
	pretty("hdd", hdd);
	pretty("ssd", ssd);
	pretty("all", all);
	pretty("cpu", cpu);
	console.log('\n');
	if (key === 'last_dumps') {
		servers
			.filter(s => s.price < 24)
			.filter(s => s.cpu_benchmark >= 9100)
			.filter(s => s.ram > 16)
			.map(calc)
			.map(s =>
`key: ${s.key}
cpu: ${s.cpu_benchmark} ${s.cpu}
ram: ${s.ram}Gb
hdd: ${s.hdd_size}
allb: ${s.allb}
cpub: ${s.cpub}
ramb: ${s.ramb}
hddb: ${s.hddb}
price: ${s.price}
reduce: ${s.next_reduce_hr}
${s.description.join('\n')}
https://www.hetzner.com/sb
https://api1.nirhub.ru/hetzner.txt
`			)
			.forEach(s => {
				bot.sendMessage('@gongo_hezner', s);
			})
	}
}
function pretty (key, servers) {
	console.log(`${key}:`);
	console.table(
		servers.map(s => ({
			...s,
			cpub: Math.round(s.c*100*10)/10,
			hddb: Math.round(s.h*100*10)/10,
			ramb: Math.round(s.r*100*10)/10,
			allb: Math.round(s.a*100*10)/10,
		}))
		.map(s => ({
			key: s.key,
			allb: s.allb>=40 ? md.format(`${s.allb}##`) : s.allb,
			cpub: s.cpub>=90 ? md.format(`${s.cpub}##`) : s.cpub,
			hddb: s.hddb>=90 ? md.format(`${s.hddb}##`) : s.hddb,
			ramb: s.ramb>=90 ? md.format(`${s.ramb}##`) : s.ramb,
			hdd: s.hdd_size,
			cpu: s.cpu,
			bench: s.cpu_benchmark,
			price: Math.round(s.price*10)/10,
			ram: s.ram,
			next: s.next_reduce_hr,
			description: s.description.join(' ')
		}))
	);
	console.log("");
}

print("all_dumps", all_dumps);
print("last_dumps", last_dump);
