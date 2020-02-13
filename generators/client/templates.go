package client

var templates = map[string]string{
	"action":        "e3sgR2VuZXJhdGVkV2FybmluZyB9fQoKcGFja2FnZSB7eyAuUGFja2FnZSB9fQoKaW1wb3J0ICgKCSJjb250ZXh0IgoJImVuY29kaW5nL2pzb24iCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL3Byb3RvY29sIgoJcnBjY2xpZW50ICJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvcHJvdmlkZXJzL2FnZW50L21jb3JwYy9jbGllbnQiCikKCi8vIHt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVxdWVzdG9yIHBlcmZvcm1zIGEgUlBDIHJlcXVlc3QgdG8ge3sgLkFnZW50TmFtZSB8IFRvTG93ZXIgfX0je3sgLkFjdGlvbk5hbWUgfCBUb0xvd2VyIH19CnR5cGUge3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXF1ZXN0b3Igc3RydWN0IHsKCXIgICAgKnJlcXVlc3RvcgoJb3V0YyBjaGFuICp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dAp9CgovLyB7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dCBpcyB0aGUgb3V0cHV0IGZyb20gdGhlIHt7IC5BY3Rpb25OYW1lIHwgVG9Mb3dlciB9fSBhY3Rpb24KdHlwZSB7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fU91dHB1dCBzdHJ1Y3QgewoJZGV0YWlscyAqUmVzdWx0RGV0YWlscwoJcmVwbHkgICBtYXBbc3RyaW5nXWludGVyZmFjZXt9Cn0KCi8vIHt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVzdWx0IGlzIHRoZSByZXN1bHQgZnJvbSBhIHt7IC5BY3Rpb25OYW1lIHwgVG9Mb3dlciB9fSBhY3Rpb24KdHlwZSB7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlc3VsdCBzdHJ1Y3QgewoJc3RhdHMgICAqcnBjY2xpZW50LlN0YXRzCglvdXRwdXRzIFtdKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0Cn0KCi8vIFN0YXRzIGlzIHRoZSBycGMgcmVxdWVzdCBzdGF0cwpmdW5jIChkICp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlc3VsdCkgU3RhdHMoKSBTdGF0cyB7CglyZXR1cm4gZC5zdGF0cwp9CgovLyBSZXN1bHREZXRhaWxzIGlzIHRoZSBkZXRhaWxzIGFib3V0IHRoZSByZXF1ZXN0CmZ1bmMgKGQgKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0KSBSZXN1bHREZXRhaWxzKCkgKlJlc3VsdERldGFpbHMgewoJcmV0dXJuIGQuZGV0YWlscwp9CgovLyBIYXNoTWFwIGlzIHRoZSByYXcgb3V0cHV0IGRhdGEKZnVuYyAoZCAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQpIEhhc2hNYXAoKSBtYXBbc3RyaW5nXWludGVyZmFjZXt9IHsKCXJldHVybiBkLnJlcGx5Cn0KCi8vIEpTT04gaXMgdGhlIEpTT04gcmVwcmVzZW50YXRpb24gb2YgdGhlIG91dHB1dCBkYXRhCmZ1bmMgKGQgKnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0KSBKU09OKCkgKFtdYnl0ZSwgZXJyb3IpIHsKCXJldHVybiBqc29uLk1hcnNoYWwoZC5yZXBseSkKfQoKLy8gRG8gcGVyZm9ybXMgdGhlIHJlcXVlc3QKZnVuYyAoZCAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXF1ZXN0b3IpIERvKGN0eCBjb250ZXh0LkNvbnRleHQpICgqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1SZXN1bHQsIGVycm9yKSB7CglkcmVzIDo9ICZ7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlc3VsdHt9CgoJaGFuZGxlciA6PSBmdW5jKHByIHByb3RvY29sLlJlcGx5LCByICpycGNjbGllbnQuUlBDUmVwbHkpIHsKCQlvdXRwdXQgOj0gJnt7IC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0ewoJCQlyZXBseTogbWFrZShtYXBbc3RyaW5nXWludGVyZmFjZXt9KSwKCQkJZGV0YWlsczogJlJlc3VsdERldGFpbHN7CgkJCQlzZW5kZXI6ICBwci5TZW5kZXJJRCgpLAoJCQkJY29kZTogICAgaW50KHIuU3RhdHVzY29kZSksCgkJCQltZXNzYWdlOiByLlN0YXR1c21zZywKCQkJCXRzOiAgICAgIHByLlRpbWUoKSwKCQkJfSwKCQl9CgoJCWVyciA6PSBqc29uLlVubWFyc2hhbChyLkRhdGEsICZvdXRwdXQucmVwbHkpCgkJaWYgZXJyICE9IG5pbCB7CgkJCWQuci5jbGllbnQuZXJyb3JmKCJDb3VsZCBub3QgZGVjb2RlIHJlcGx5IGZyb20gJXM6ICVzIiwgcHIuU2VuZGVySUQoKSwgZXJyKQoJCX0KCgkJaWYgZC5vdXRjICE9IG5pbCB7CgkJCWQub3V0YyA8LSBvdXRwdXQKCQkJcmV0dXJuCgkJfQoKCQlkcmVzLm91dHB1dHMgPSBhcHBlbmQoZHJlcy5vdXRwdXRzLCBvdXRwdXQpCgl9CgoJcmVzLCBlcnIgOj0gZC5yLmRvKGN0eCwgaGFuZGxlcikKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGVycgoJfQoKCWRyZXMuc3RhdHMgPSByZXMKCglyZXR1cm4gZHJlcywgbmlsCn0KCi8vIEVhY2hPdXRwdXQgaXRlcmF0ZXMgb3ZlciBhbGwgcmVzdWx0cyByZWNlaXZlZApmdW5jIChkICp7eyAuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlc3VsdCkgRWFjaE91dHB1dChoIGZ1bmMociAqe3sgLkFjdGlvbk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1PdXRwdXQpKSB7Cglmb3IgXywgcmVzcCA6PSByYW5nZSBkLm91dHB1dHMgewoJCWgocmVzcCkKCX0KfQoKe3sgcmFuZ2UgJG5hbWUsICRpbnB1dCA6PSAuT3B0aW9uYWxJbnB1dHMgfX0KLy8ge3sgJG5hbWUgfCBTbmFrZVRvQ2FtZWwgfX0gaXMgYW4gb3B0aW9uYWwgaW5wdXQgdG8gdGhlIHt7ICQuQWN0aW9uTmFtZSB8IFRvTG93ZXIgfX0gYWN0aW9uCi8vCi8vIERlc2NyaXB0aW9uOiB7eyAkaW5wdXQuRGVzY3JpcHRpb24gfX0KZnVuYyAoZCAqe3sgJC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVxdWVzdG9yKSB7eyAkbmFtZSB8IFNuYWtlVG9DYW1lbCB9fSh2IHt7IENob3JpYVR5cGVUb0dvVHlwZSAkaW5wdXQuVHlwZSB9fSkgKnt7ICQuQWN0aW9uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlcXVlc3RvciB7CglkLnIuYXJnc1sie3sgJG5hbWUgfCBUb0xvd2VyIH19Il0gPSB2CgoJcmV0dXJuIGQKfQp7eyBlbmQgfX0Ke3sgcmFuZ2UgJG5hbWUsICRvdXRwdXQgOj0gLk91dHB1dHMgfX0KLy8ge3sgJG5hbWUgfCBTbmFrZVRvQ2FtZWwgfX0gaXMgdGhlIHZhbHVlIG9mIHRoZSB7eyAkbmFtZSB9fSBvdXRwdXQKLy8KLy8gRGVzY3JpcHRpb246IHt7ICRvdXRwdXQuRGVzY3JpcHRpb24gfX0KZnVuYyAoZCAqe3sgJC5BY3Rpb25OYW1lIHwgU25ha2VUb0NhbWVsIH19T3V0cHV0KSB7eyAkbmFtZSB8IFNuYWtlVG9DYW1lbCB9fSgpIHt7IENob3JpYVR5cGVUb0dvVHlwZSAkb3V0cHV0LlR5cGUgfX0gewoJdmFsIDo9IGQucmVwbHlbInt7ICRuYW1lIH19Il0KCXJldHVybiB7eyBDaG9yaWFUeXBlVG9WYWxPZlR5cGUgJG91dHB1dC5UeXBlIH19Cn0Ke3sgZW5kIH19Cg==",
	"client":        "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImVuY29kaW5nL2Jhc2U2NCIKCSJlbmNvZGluZy9qc29uIgoJImZtdCIKCSJ0aW1lIgoKCSJjb250ZXh0IgoKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvY2hvcmlhIgoJY29yZWNsaWVudCAiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL2NsaWVudC9jbGllbnQiCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL2NvbmZpZyIKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvc3J2Y2FjaGUiCglycGNjbGllbnQgImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9wcm92aWRlcnMvYWdlbnQvbWNvcnBjL2NsaWVudCIKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvcHJvdmlkZXJzL2FnZW50L21jb3JwYy9kZGwvYWdlbnQiCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL3Byb3RvY29sIgoJImdpdGh1Yi5jb20vc2lydXBzZW4vbG9ncnVzIgopCgovLyBTdGF0cyBhcmUgdGhlIHN0YXRpc3RpY3MgZm9yIGEgcmVxdWVzdAp0eXBlIFN0YXRzIGludGVyZmFjZSB7CglBZ2VudCgpIHN0cmluZwoJQWN0aW9uKCkgc3RyaW5nCglBbGwoKSBib29sCglOb1Jlc3BvbnNlRnJvbSgpIFtdc3RyaW5nCglVbmV4cGVjdGVkUmVzcG9uc2VGcm9tKCkgW11zdHJpbmcKCURpc2NvdmVyZWRDb3VudCgpIGludAoJRGlzY292ZXJlZE5vZGVzKCkgKltdc3RyaW5nCglGYWlsQ291bnQoKSBpbnQKCU9LQ291bnQoKSBpbnQKCVJlc3BvbnNlc0NvdW50KCkgaW50CglQdWJsaXNoRHVyYXRpb24oKSAodGltZS5EdXJhdGlvbiwgZXJyb3IpCglSZXF1ZXN0RHVyYXRpb24oKSAodGltZS5EdXJhdGlvbiwgZXJyb3IpCglEaXNjb3ZlcnlEdXJhdGlvbigpICh0aW1lLkR1cmF0aW9uLCBlcnJvcikKfQoKLy8gTm9kZVNvdXJjZSBkaXNjb3ZlcnMgbm9kZXMKdHlwZSBOb2RlU291cmNlIGludGVyZmFjZSB7CglSZXNldCgpCglEaXNjb3ZlcihjdHggY29udGV4dC5Db250ZXh0LCBmdyBDaG9yaWFGcmFtZXdvcmssIGZpbHRlcnMgW11GaWx0ZXJGdW5jKSAoW11zdHJpbmcsIGVycm9yKQp9CgovLyBDaG9yaWFGcmFtZXdvcmsgaXMgdGhlIENob3JpYSBmcmFtZXdvcmsKdHlwZSBDaG9yaWFGcmFtZXdvcmsgaW50ZXJmYWNlIHsKCUxvZ2dlcihzdHJpbmcpICpsb2dydXMuRW50cnkKCUNvbmZpZ3VyYXRpb24oKSAqY29uZmlnLkNvbmZpZwoJTmV3TWVzc2FnZShwYXlsb2FkIHN0cmluZywgYWdlbnQgc3RyaW5nLCBjb2xsZWN0aXZlIHN0cmluZywgbXNnVHlwZSBzdHJpbmcsIHJlcXVlc3QgKmNob3JpYS5NZXNzYWdlKSAobXNnICpjaG9yaWEuTWVzc2FnZSwgZXJyIGVycm9yKQoJTmV3UmVwbHlGcm9tVHJhbnNwb3J0SlNPTihwYXlsb2FkIFtdYnl0ZSwgc2tpcHZhbGlkYXRlIGJvb2wpIChtc2cgcHJvdG9jb2wuUmVwbHksIGVyciBlcnJvcikKCU5ld1RyYW5zcG9ydEZyb21KU09OKGRhdGEgc3RyaW5nKSAobWVzc2FnZSBwcm90b2NvbC5UcmFuc3BvcnRNZXNzYWdlLCBlcnIgZXJyb3IpCglNaWRkbGV3YXJlU2VydmVycygpIChzZXJ2ZXJzIHNydmNhY2hlLlNlcnZlcnMsIGVyciBlcnJvcikKCU5ld0Nvbm5lY3RvcihjdHggY29udGV4dC5Db250ZXh0LCBzZXJ2ZXJzIGZ1bmMoKSAoc3J2Y2FjaGUuU2VydmVycywgZXJyb3IpLCBuYW1lIHN0cmluZywgbG9nZ2VyICpsb2dydXMuRW50cnkpIChjb25uIGNob3JpYS5Db25uZWN0b3IsIGVyciBlcnJvcikKCU5ld1JlcXVlc3RJRCgpIChzdHJpbmcsIGVycm9yKQoJQ2VydG5hbWUoKSBzdHJpbmcKfQoKLy8gRmlsdGVyRnVuYyBjYW4gZ2VuZXJhdGUgYSBDaG9yaWEgZmlsdGVyCnR5cGUgRmlsdGVyRnVuYyBmdW5jKGYgKnByb3RvY29sLkZpbHRlcikgZXJyb3IKCi8vIHt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB0byB0aGUge3sgLkRETC5NZXRhZGF0YS5OYW1lIH19IGFnZW50CnR5cGUge3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHN0cnVjdCB7CglmdyAgICAgICAgICAgQ2hvcmlhRnJhbWV3b3JrCgljZmcgICAgICAgICAgICpjb25maWcuQ29uZmlnCglkZGwgICAgICAgICAgICphZ2VudC5EREwKCW5zICAgICAgICAgICAgTm9kZVNvdXJjZQoJY2xpZW50T3B0cyAgICAqaW5pdE9wdGlvbnMKCWNsaWVudFJQQ09wdHMgW11ycGNjbGllbnQuUmVxdWVzdE9wdGlvbgoJZmlsdGVycyAgICAgICBbXUZpbHRlckZ1bmMKfQoKLy8gTWV0YWRhdGEgaXMgdGhlIGFnZW50IG1ldGFkYXRhCnR5cGUgTWV0YWRhdGEgc3RydWN0IHsKCUxpY2Vuc2UgICAgIHN0cmluZyBganNvbjoibGljZW5zZSJgCglBdXRob3IgICAgICBzdHJpbmcgYGpzb246ImF1dGhvciJgCglUaW1lb3V0ICAgICBpbnQgICAgYGpzb246InRpbWVvdXQiYAoJTmFtZSAgICAgICAgc3RyaW5nIGBqc29uOiJuYW1lImAKCVZlcnNpb24gICAgIHN0cmluZyBganNvbjoidmVyc2lvbiJgCglVUkwgICAgICAgICBzdHJpbmcgYGpzb246InVybCJgCglEZXNjcmlwdGlvbiBzdHJpbmcgYGpzb246ImRlc2NyaXB0aW9uImAKfQoKLy8gTXVzdCBjcmVhdGUgYSBuZXcgY2xpZW50IGFuZCBwYW5pY3Mgb24gZXJyb3IKZnVuYyBNdXN0KG9wdHMgLi4uSW5pdGlhbGl6YXRpb25PcHRpb24pIChjbGllbnQgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgewoJYywgZXJyIDo9IE5ldyhvcHRzLi4uKQoJaWYgZXJyICE9IG5pbCB7CgkJcGFuaWMoZXJyKQoJfQoKCXJldHVybiBjCn0KCi8vIE5ldyBjcmVhdGVzIGEgbmV3IGNsaWVudCB0byB0aGUge3sgLkRETC5NZXRhZGF0YS5OYW1lIH19IGFnZW50CmZ1bmMgTmV3KG9wdHMgLi4uSW5pdGlhbGl6YXRpb25PcHRpb24pIChjbGllbnQgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCwgZXJyIGVycm9yKSB7CgljIDo9ICZ7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnR7CgkJZGRsOiAgICAgICAgICAgJmFnZW50LkRETHt9LAoJCWNsaWVudFJQQ09wdHM6IFtdcnBjY2xpZW50LlJlcXVlc3RPcHRpb257fSwKCQlmaWx0ZXJzOiAgICAgICBbXUZpbHRlckZ1bmN7CgkJICAgIEZpbHRlckZ1bmMoY29yZWNsaWVudC5BZ2VudEZpbHRlcigie3sgLkRETC5NZXRhZGF0YS5OYW1lIH19IikpLAoJCX0sCgkJY2xpZW50T3B0czogJmluaXRPcHRpb25zewoJCQljZmdGaWxlOiBjaG9yaWEuVXNlckNvbmZpZygpLAoJCX0sCgl9CgoJZm9yIF8sIG9wdCA6PSByYW5nZSBvcHRzIHsKCQlvcHQoYy5jbGllbnRPcHRzKQoJfQoKCWlmIGMuY2xpZW50T3B0cy5ucyA9PSBuaWwgewoJCWMuY2xpZW50T3B0cy5ucyA9ICZCcm9hZGNhc3ROU3t9Cgl9CgljLm5zID0gYy5jbGllbnRPcHRzLm5zCgoJYy5mdywgZXJyID0gY2hvcmlhLk5ldyhjLmNsaWVudE9wdHMuY2ZnRmlsZSkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoImNvdWxkIG5vdCBpbml0aWFsaXplIENob3JpYTogJXMiLCBlcnIpCgl9CgoJYy5jZmcgPSBjLmZ3LkNvbmZpZ3VyYXRpb24oKQoKCWlmIGMuY2xpZW50T3B0cy5sb2dnZXIgPT0gbmlsIHsKCQljLmNsaWVudE9wdHMubG9nZ2VyID0gYy5mdy5Mb2dnZXIoInt7IC5EREwuTWV0YWRhdGEuTmFtZSB9fSIpCgl9CgoJZGRsaiwgZXJyIDo9IGJhc2U2NC5TdGRFbmNvZGluZy5EZWNvZGVTdHJpbmcocmF3RERMKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIG5pbCwgZm10LkVycm9yZigiY291bGQgbm90IHBhcnNlIGVtYmVkZGVkIERETDogJXMiLCBlcnIpCgl9CgoJZXJyID0ganNvbi5Vbm1hcnNoYWwoZGRsaiwgYy5kZGwpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBmbXQuRXJyb3JmKCJjb3VsZCBub3QgcGFyc2UgZW1iZWRkZWQgRERMOiAlcyIsIGVycikKCX0KCglyZXR1cm4gYywgbmlsCn0KCi8vIEFnZW50TWV0YWRhdGEgaXMgdGhlIGFnZW50IG1ldGFkYXRhIHRoaXMgY2xpZW50IHN1cHBvcnRzCmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgQWdlbnRNZXRhZGF0YSgpICpNZXRhZGF0YSB7CglyZXR1cm4gJk1ldGFkYXRhewoJCUxpY2Vuc2U6ICAgICBwLmRkbC5NZXRhZGF0YS5MaWNlbnNlLAoJCUF1dGhvcjogICAgICBwLmRkbC5NZXRhZGF0YS5BdXRob3IsCgkJVGltZW91dDogICAgIHAuZGRsLk1ldGFkYXRhLlRpbWVvdXQsCgkJTmFtZTogICAgICAgIHAuZGRsLk1ldGFkYXRhLk5hbWUsCgkJVmVyc2lvbjogICAgIHAuZGRsLk1ldGFkYXRhLlZlcnNpb24sCgkJVVJMOiAgICAgICAgIHAuZGRsLk1ldGFkYXRhLlVSTCwKCQlEZXNjcmlwdGlvbjogcC5kZGwuTWV0YWRhdGEuRGVzY3JpcHRpb24sCgl9Cn0KCnt7IHJhbmdlICRpLCAkYWN0aW9uIDo9IC5EREwuQWN0aW9ucyB9fQovLyB7eyAkYWN0aW9uLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX0gcGVyZm9ybXMgdGhlIHt7ICRhY3Rpb24uTmFtZSB8IFRvTG93ZXIgfX0gYWN0aW9uCi8vCi8vIERlc2NyaXB0aW9uOiB7eyAkYWN0aW9uLkRlc2NyaXB0aW9uIH19Cnt7LSBpZiBDaG9yaWFSZXF1aXJlZElucHV0cyAkYWN0aW9uIH19Ci8vCi8vIFJlcXVpcmVkIElucHV0czoKe3stIHJhbmdlICRuYW1lLCAkaW5wdXQgOj0gQ2hvcmlhUmVxdWlyZWRJbnB1dHMgJGFjdGlvbiB9fQovLyAgICAtIHt7ICRuYW1lIH19ICh7eyAkaW5wdXQuVHlwZSB8IENob3JpYVR5cGVUb0dvVHlwZSB9fSkgLSB7eyAkaW5wdXQuRGVzY3JpcHRpb24gfX0Ke3stIGVuZCB9fQp7ey0gZW5kIH19Cnt7LSBpZiBDaG9yaWFPcHRpb25hbElucHV0cyAkYWN0aW9uIH19Ci8vCi8vIE9wdGlvbmFsIElucHV0czoKe3stIHJhbmdlICRuYW1lLCAkaW5wdXQgOj0gQ2hvcmlhT3B0aW9uYWxJbnB1dHMgJGFjdGlvbiB9fQovLyAgICAtIHt7ICRuYW1lIH19ICh7eyAkaW5wdXQuVHlwZSB8IENob3JpYVR5cGVUb0dvVHlwZSB9fSkgLSB7eyAkaW5wdXQuRGVzY3JpcHRpb24gfX0Ke3stIGVuZCB9fQp7ey0gZW5kIH19CmZ1bmMgKHAgKnt7ICQuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIHt7ICRhY3Rpb24uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fSh7eyAkYWN0aW9uIHwgQ2hvcmlhUmVxdWlyZWRJbnB1dHNUb0Z1bmNBcmdzIH19KSAqe3sgJGFjdGlvbi5OYW1lIHwgU25ha2VUb0NhbWVsIH19UmVxdWVzdG9yIHsKCWQgOj0gJnt7ICRhY3Rpb24uTmFtZSB8IFNuYWtlVG9DYW1lbCB9fVJlcXVlc3RvcnsKCQlvdXRjOiBuaWwsCgkJcjogJnJlcXVlc3RvcnsKCQkJYXJnczogICBtYXBbc3RyaW5nXWludGVyZmFjZXt9ewp7ey0gcmFuZ2UgJG5hbWUsICRpbnB1dCA6PSBDaG9yaWFSZXF1aXJlZElucHV0cyAkYWN0aW9uIH19CgkJCQkie3sgJG5hbWUgfX0iOiB7eyAkbmFtZSB9fUksCnt7LSBlbmQgfX0KCQkJfSwKCQkJYWN0aW9uOiAie3sgJGFjdGlvbi5OYW1lIHwgVG9Mb3dlciB9fSIsCgkJCWNsaWVudDogcCwKCQl9LAoJfQoKCXJldHVybiBkCn0Ke3sgZW5kIH19Cg==",
	"ddl":           "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19Cgp2YXIgcmF3RERMID0gInt7IC5SYXdEREwgfCBCYXNlNjRFbmNvZGUgfX0i",
	"discover":      "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImNvbnRleHQiCgkic3luYyIKCSJ0aW1lIgoKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvY2xpZW50L2Rpc2NvdmVyeS9icm9hZGNhc3QiCgkiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL3Byb3RvY29sIgopCgovLyBCcm9hZGNhc3ROUyBpcyBhIE5vZGVTb3VyY2UgdGhhdCB1c2VzIHRoZSBDaG9yaWEgbmV0d29yayBicm9hZGNhc3QgbWV0aG9kIHRvIGRpc2NvdmVyIG5vZGVzCnR5cGUgQnJvYWRjYXN0TlMgc3RydWN0IHsKCW5vZGVDYWNoZSBbXXN0cmluZwoJZiAgICAgICAgICpwcm90b2NvbC5GaWx0ZXIKCglzeW5jLk11dGV4Cn0KCi8vIFJlc2V0IHJlc2V0cyB0aGUgaW50ZXJuYWwgbm9kZSBjYWNoZQpmdW5jIChiICpCcm9hZGNhc3ROUykgUmVzZXQoKSB7CgliLkxvY2soKQoJZGVmZXIgYi5VbmxvY2soKQoKCWIubm9kZUNhY2hlID0gW11zdHJpbmd7fQp9CgovLyBEaXNjb3ZlciBwZXJmb3JtcyB0aGUgZGlzY292ZXJ5IG9mIG5vZGVzIGFnYWluc3QgdGhlIENob3JpYSBOZXR3b3JrCmZ1bmMgKGIgKkJyb2FkY2FzdE5TKSBEaXNjb3ZlcihjdHggY29udGV4dC5Db250ZXh0LCBmdyBDaG9yaWFGcmFtZXdvcmssIGZpbHRlcnMgW11GaWx0ZXJGdW5jKSAoW11zdHJpbmcsIGVycm9yKSB7CgliLkxvY2soKQoJZGVmZXIgYi5VbmxvY2soKQoKCWNvcGllciA6PSBmdW5jKCkgW11zdHJpbmcgewoJCW91dCA6PSBtYWtlKFtdc3RyaW5nLCBsZW4oYi5ub2RlQ2FjaGUpKQoJCWZvciBpLCBuIDo9IHJhbmdlIGIubm9kZUNhY2hlIHsKCQkJb3V0W2ldID0gbgoJCX0KCgkJcmV0dXJuIG91dAoJfQoKCWlmICEoYi5ub2RlQ2FjaGUgPT0gbmlsIHx8IGxlbihiLm5vZGVDYWNoZSkgPT0gMCkgewoJCXJldHVybiBjb3BpZXIoKSwgbmlsCgl9CgoJZXJyIDo9IGIucGFyc2VGaWx0ZXJzKGZpbHRlcnMpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBlcnIKCX0KCglpZiBiLm5vZGVDYWNoZSA9PSBuaWwgewoJCWIubm9kZUNhY2hlID0gW11zdHJpbmd7fQoJfQoKCWNmZyA6PSBmdy5Db25maWd1cmF0aW9uKCkKCW5vZGVzLCBlcnIgOj0gYnJvYWRjYXN0Lk5ldyhmdykuRGlzY292ZXIoY3R4LCBicm9hZGNhc3QuRmlsdGVyKGIuZiksIGJyb2FkY2FzdC5UaW1lb3V0KHRpbWUuU2Vjb25kKnRpbWUuRHVyYXRpb24oY2ZnLkRpc2NvdmVyeVRpbWVvdXQpKSkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBbXXN0cmluZ3t9LCBlcnIKCX0KCgliLm5vZGVDYWNoZSA9IG5vZGVzCgoJcmV0dXJuIGNvcGllcigpLCBuaWwKfQoKZnVuYyAoYiAqQnJvYWRjYXN0TlMpIHBhcnNlRmlsdGVycyhmcyBbXUZpbHRlckZ1bmMpIGVycm9yIHsKCWIuZiA9IHByb3RvY29sLk5ld0ZpbHRlcigpCgoJZm9yIF8sIGYgOj0gcmFuZ2UgZnMgewoJCWVyciA6PSBmKGIuZikKCQlpZiBlcnIgIT0gbmlsIHsKCQkJcmV0dXJuIGVycgoJCX0KCX0KCglyZXR1cm4gbmlsCn0K",
	"doc":           "e3sgR2VuZXJhdGVkV2FybmluZyB9fQoKLy8gUGFja2FnZSB7eyAuUGFja2FnZSB9fSBpcyBhbiBBUEkgY2xpZW50IHRvIHRoZSBDaG9yaWEge3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgQ2FwaXRhbGl6ZSB9fSBhZ2VudCBWZXJzaW9uIHt7IC5EREwuTWV0YWRhdGEuVmVyc2lvbiB9fS4KLy8KLy8gQWN0aW9uczoKe3stIHJhbmdlICRpLCAkYWN0aW9uIDo9IC5EREwuQWN0aW9ucyB9fQovLyAgICoge3sgJGFjdGlvbi5OYW1lIHwgU25ha2VUb0NhbWVsIH19IC0ge3sgJGFjdGlvbi5EZXNjcmlwdGlvbiAtfX0Ke3sgZW5kIH19CnBhY2thZ2Uge3sgLlBhY2thZ2UgfX0KCg==",
	"initoptions":   "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImdpdGh1Yi5jb20vc2lydXBzZW4vbG9ncnVzIgopCgp0eXBlIGluaXRPcHRpb25zIHN0cnVjdCB7CgljZmdGaWxlIHN0cmluZwoJbG9nZ2VyICAqbG9ncnVzLkVudHJ5CglucyBOb2RlU291cmNlCn0KCi8vIEluaXRpYWxpemF0aW9uT3B0aW9uIGlzIGFuIG9wdGlvbmFsIHNldHRpbmcgdXNlZCB0byBpbml0aWFsaXplIHRoZSBjbGllbnQKdHlwZSBJbml0aWFsaXphdGlvbk9wdGlvbiBmdW5jKG9wdHMgKmluaXRPcHRpb25zKQoKLy8gQ29uZmlnRmlsZSBzZXRzIHRoZSBjb25maWcgZmlsZSB0byB1c2UsIHdoZW4gbm90IHNldCB3aWxsIHVzZSB0aGUgdXNlciBkZWZhdWx0CmZ1bmMgQ29uZmlnRmlsZShmIHN0cmluZykgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLmNmZ0ZpbGUgPSBmCgl9Cn0KCi8vIExvZ2dlciBzZXRzIHRoZSBsb2dnZXIgdG8gdXNlIGVsc2Ugb25lIGlzIG1hZGUgdmlhIHRoZSBDaG9yaWEgZnJhbWV3b3JrCmZ1bmMgTG9nZ2VyKGwgKmxvZ3J1cy5FbnRyeSkgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLmxvZ2dlciA9IGwKCX0KfQoKLy8gRGlzY292ZXJ5IHNldHMgdGhlIE5vZGVTb3VyY2UgdG8gdXNlIHdoZW4gZmluZGluZyBub2RlcyB0byBtYW5hZ2UKZnVuYyBEaXNjb3ZlcnkobnMgTm9kZVNvdXJjZSkgSW5pdGlhbGl6YXRpb25PcHRpb24gewoJcmV0dXJuIGZ1bmMobyAqaW5pdE9wdGlvbnMpIHsKCQlvLm5zID0gbnMKCX0KfQo=",
	"logging":       "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgpmdW5jIChjICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIGRlYnVnZihtc2cgc3RyaW5nLCBhIC4uLmludGVyZmFjZXt9KSB7CgljLmNsaWVudE9wdHMubG9nZ2VyLkRlYnVnZihtc2csIGEuLi4pCn0KCmZ1bmMgKGMgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgaW5mb2YobXNnIHN0cmluZywgYSAuLi5pbnRlcmZhY2V7fSkgewoJYy5jbGllbnRPcHRzLmxvZ2dlci5JbmZvZihtc2csIGEuLi4pCn0KCmZ1bmMgKGMgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgd2FybmYobXNnIHN0cmluZywgYSAuLi5pbnRlcmZhY2V7fSkgewoJYy5jbGllbnRPcHRzLmxvZ2dlci5XYXJuZihtc2csIGEuLi4pCn0KCmZ1bmMgKGMgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgZXJyb3JmKG1zZyBzdHJpbmcsIGEgLi4uaW50ZXJmYWNle30pIHsKCWMuY2xpZW50T3B0cy5sb2dnZXIuRXJyb3JmKG1zZywgYS4uLikKfQo=",
	"requestor":     "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJImNvbnRleHQiCgkiZm10IgoKCSJnaXRodWIuY29tL2Nob3JpYS1pby9nby1jaG9yaWEvcHJvdG9jb2wiCglycGNjbGllbnQgImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9wcm92aWRlcnMvYWdlbnQvbWNvcnBjL2NsaWVudCIKKQoKLy8gcmVxdWVzdG9yIGlzIGEgZ2VuZXJpYyByZXF1ZXN0IGhhbmRsZXIKdHlwZSByZXF1ZXN0b3Igc3RydWN0IHsKCWNsaWVudCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50CglhY3Rpb24gc3RyaW5nCglhcmdzICAgbWFwW3N0cmluZ11pbnRlcmZhY2V7fQp9CgovLyBkbyBwZXJmb3JtcyB0aGUgcmVxdWVzdApmdW5jIChyICpyZXF1ZXN0b3IpIGRvKGN0eCBjb250ZXh0LkNvbnRleHQsIGhhbmRsZXIgZnVuYyhwciBwcm90b2NvbC5SZXBseSwgciAqcnBjY2xpZW50LlJQQ1JlcGx5KSkgKCpycGNjbGllbnQuU3RhdHMsIGVycm9yKSB7CglyLmNsaWVudC5pbmZvZigiU3RhcnRpbmcgZGlzY292ZXJ5IikKCXRhcmdldHMsIGVyciA6PSByLmNsaWVudC5ucy5EaXNjb3ZlcihjdHgsIHIuY2xpZW50LmZ3LCByLmNsaWVudC5maWx0ZXJzKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIG5pbCwgZXJyCgl9CgoJaWYgbGVuKHRhcmdldHMpID09IDAgewoJCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoIm5vIG5vZGVzIHdlcmUgZGlzY292ZXJlZCIpCgl9CglyLmNsaWVudC5pbmZvZigiRGlzY292ZXJlZCAlZCBub2RlcyIsIGxlbih0YXJnZXRzKSkKCglhZ2VudCwgZXJyIDo9IHJwY2NsaWVudC5OZXcoci5jbGllbnQuZncsIHIuY2xpZW50LmRkbC5NZXRhZGF0YS5OYW1lLCBycGNjbGllbnQuRERMKHIuY2xpZW50LmRkbCkpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBmbXQuRXJyb3JmKCJjb3VsZCBub3QgY3JlYXRlIGNsaWVudDogJXMiLCBlcnIpCgl9CgoJb3B0cyA6PSBbXXJwY2NsaWVudC5SZXF1ZXN0T3B0aW9ue3JwY2NsaWVudC5UYXJnZXRzKHRhcmdldHMpfQoJZm9yIF8sIG9wdCA6PSByYW5nZSByLmNsaWVudC5jbGllbnRSUENPcHRzIHsKCQlvcHRzID0gYXBwZW5kKG9wdHMsIG9wdCkKCX0KCW9wdHMgPSBhcHBlbmQob3B0cywgcnBjY2xpZW50LlJlcGx5SGFuZGxlcihoYW5kbGVyKSkKCglyLmNsaWVudC5pbmZvZigiSW52b2tpbmcgJXMjJXMgYWN0aW9uIHdpdGggJSN2Iiwgci5jbGllbnQuZGRsLk1ldGFkYXRhLk5hbWUsIHIuYWN0aW9uLCByLmFyZ3MpCgoJcmVzLCBlcnIgOj0gYWdlbnQuRG8oY3R4LCByLmFjdGlvbiwgci5hcmdzLCBvcHRzLi4uKQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuIG5pbCwgZm10LkVycm9yZigiY291bGQgbm90IHBlcmZvcm0gZGlzYWJsZSByZXF1ZXN0OiAlcyIsIGVycikKCX0KCglyZXR1cm4gcmVzLlN0YXRzKCksIG5pbAp9Cg==",
	"resultdetails": "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJInRpbWUiCgoJImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9wcm92aWRlcnMvYWdlbnQvbWNvcnBjIgopCgovLyBSZXN1bHREZXRhaWxzIGlzIHRoZSBkZXRhaWxzIGFib3V0IGEgcmVzdWx0CnR5cGUgUmVzdWx0RGV0YWlscyBzdHJ1Y3QgewoJc2VuZGVyICBzdHJpbmcKCWNvZGUgICAgaW50CgltZXNzYWdlIHN0cmluZwoJdHMgICAgICB0aW1lLlRpbWUKfQoKLy8gU3RhdHVzQ29kZSBpcyBhIHJlcGx5IHN0YXR1cyBhcyBkZWZpbmVkIGJ5IE1Db2xsZWN0aXZlIFNpbXBsZVJQQyAtIGludGVnZXJzIDAgdG8gNQovLwovLyBTZWUgdGhlIGNvbnN0YW50cyBPSywgUlBDQWJvcnRlZCwgVW5rbm93blJQQ0FjdGlvbiwgTWlzc2luZ1JQQ0RhdGEsIEludmFsaWRSUENEYXRhIGFuZCBVbmtub3duUlBDRXJyb3IKdHlwZSBTdGF0dXNDb2RlIHVpbnQ4Cgpjb25zdCAoCgkvLyBPSyBpcyB0aGUgcmVwbHkgc3RhdHVzIHdoZW4gYWxsIHdvcmtlZAoJT0sgPSBTdGF0dXNDb2RlKGlvdGEpCgoJLy8gQWJvcnRlZCBpcyBzdGF0dXMgZm9yIHdoZW4gdGhlIGFjdGlvbiBjb3VsZCBub3QgcnVuLCBtb3N0IGZhaWx1cmVzIGluIGFuIGFjdGlvbiBzaG91bGQgc2V0IHRoaXMKCUFib3J0ZWQKCgkvLyBVbmtub3duQWN0aW9uIGlzIHRoZSBzdGF0dXMgZm9yIHVua25vd24gYWN0aW9ucyByZXF1ZXN0ZWQKCVVua25vd25BY3Rpb24KCgkvLyBNaXNzaW5nRGF0YSBpcyB0aGUgc3RhdHVzIGZvciBtaXNzaW5nIGlucHV0IGRhdGEKCU1pc3NpbmdEYXRhCgoJLy8gSW52YWxpZERhdGEgaXMgdGhlIHN0YXR1cyBmb3IgaW52YWxpZCBpbnB1dCBkYXRhCglJbnZhbGlkRGF0YQoKCS8vIFVua25vd25FcnJvciBpcyB0aGUgc3RhdHVzIGdlbmVyYWwgZmFpbHVyZXMgaW4gYWdlbnRzIHNob3VsZCBzZXQgd2hlbiB0aGluZ3MgZ28gYmFkCglVbmtub3duRXJyb3IKKQoKLy8gU2VuZGVyIGlzIHRoZSBpZGVudGl0eSBvZiB0aGUgcmVtb3RlIHRoYXQgcHJvZHVjZWQgdGhlIG1lc3NhZ2UKZnVuYyAoZCAqUmVzdWx0RGV0YWlscykgU2VuZGVyKCkgc3RyaW5nIHsKCXJldHVybiBkLnNlbmRlcgp9CgovLyBPSyBkZXRlcm1pbmVzIGlmIHRoZSByZXF1ZXN0IHdhcyBzdWNjZXNzZnVsbApmdW5jIChkICpSZXN1bHREZXRhaWxzKSBPSygpIGJvb2wgewoJcmV0dXJuIG1jb3JwYy5TdGF0dXNDb2RlKGQuY29kZSkgPT0gbWNvcnBjLk9LCn0KCi8vIFN0YXR1c01lc3NhZ2UgaXMgdGhlIHN0YXR1cyBtZXNzYWdlIHByb2R1Y2VkIGJ5IHRoZSByZW1vdGUKZnVuYyAoZCAqUmVzdWx0RGV0YWlscykgU3RhdHVzTWVzc2FnZSgpIHN0cmluZyB7CglyZXR1cm4gZC5tZXNzYWdlCn0KCi8vIFN0YXR1c0NvZGUgaXMgdGhlIHN0YXR1cyBjb2RlIHByb2R1Y2VkIGJ5IHRoZSByZW1vdGUKZnVuYyAoZCAqUmVzdWx0RGV0YWlscykgU3RhdHVzQ29kZSgpIFN0YXR1c0NvZGUgewoJcmV0dXJuIFN0YXR1c0NvZGUoZC5jb2RlKQp9Cg==",
	"rpcoptions":    "Ly8gZ2VuZXJhdGVkIGNvZGU7IERPIE5PVCBFRElUCgpwYWNrYWdlIHt7IC5QYWNrYWdlIH19CgppbXBvcnQgKAoJInRpbWUiCgoJY29yZWNsaWVudCAiZ2l0aHViLmNvbS9jaG9yaWEtaW8vZ28tY2hvcmlhL2NsaWVudC9jbGllbnQiCglycGNjbGllbnQgImdpdGh1Yi5jb20vY2hvcmlhLWlvL2dvLWNob3JpYS9wcm92aWRlcnMvYWdlbnQvbWNvcnBjL2NsaWVudCIKKQoKLy8gT3B0aW9uUmVzZXQgcmVzZXRzIHRoZSBjbGllbnQgb3B0aW9ucyB0byB1c2UgYWNyb3NzIHJlcXVlc3RzIHRvIGFuIGVtcHR5IGxpc3QKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25SZXNldCgpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5jbGllbnRSUENPcHRzID0gW11ycGNjbGllbnQuUmVxdWVzdE9wdGlvbnt9CglwLm5zID0gcC5jbGllbnRPcHRzLm5zCglwLmZpbHRlcnMgPSBbXUZpbHRlckZ1bmN7CgkgICAgRmlsdGVyRnVuYyhjb3JlY2xpZW50LkFnZW50RmlsdGVyKCJ7eyAuRERMLk1ldGFkYXRhLk5hbWUgfX0iKSksCgl9CgoJcmV0dXJuIHAKfQoKLy8gT3B0aW9uSWRlbnRpdHlGaWx0ZXIgYWRkcyBhbiBpZGVudGl0eSBmaWx0ZXIKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25JZGVudGl0eUZpbHRlcihmIC4uLnN0cmluZykgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7Cglmb3IgXywgaSA6PSByYW5nZSBmIHsKCQlwLmZpbHRlcnMgPSBhcHBlbmQocC5maWx0ZXJzLCBGaWx0ZXJGdW5jKGNvcmVjbGllbnQuSWRlbnRpdHlGaWx0ZXIoaSkpKQoJfQoKCXAubnMuUmVzZXQoKQoKCXJldHVybiBwCn0KCi8vIE9wdGlvbkNsYXNzRmlsdGVyIGFkZHMgYSBjbGFzcyBmaWx0ZXIKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25DbGFzc0ZpbHRlcihmIC4uLnN0cmluZykgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7Cglmb3IgXywgaSA6PSByYW5nZSBmIHsKCQlwLmZpbHRlcnMgPSBhcHBlbmQocC5maWx0ZXJzLCBGaWx0ZXJGdW5jKGNvcmVjbGllbnQuQ2xhc3NGaWx0ZXIoaSkpKQoJfQoKCXAubnMuUmVzZXQoKQoKCXJldHVybiBwCn0KCi8vIE9wdGlvbkZhY3RGaWx0ZXIgYWRkcyBhIGZhY3QgZmlsdGVyCmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uRmFjdEZpbHRlcihmIC4uLnN0cmluZykgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7Cglmb3IgXywgaSA6PSByYW5nZSBmIHsKCQlwLmZpbHRlcnMgPSBhcHBlbmQocC5maWx0ZXJzLCBGaWx0ZXJGdW5jKGNvcmVjbGllbnQuRmFjdEZpbHRlcihpKSkpCgl9CgoJcC5ucy5SZXNldCgpCgoJcmV0dXJuIHAKfQoKLy8gT3B0aW9uQ29sbGVjdGl2ZSBzZXRzIHRoZSBjb2xsZWN0aXZlIHRvIHRhcmdldApmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbkNvbGxlY3RpdmUoYyBzdHJpbmcpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5jbGllbnRSUENPcHRzID0gYXBwZW5kKHAuY2xpZW50UlBDT3B0cywgcnBjY2xpZW50LkNvbGxlY3RpdmUoYykpCglyZXR1cm4gcAp9CgovLyBPcHRpb25JbkJhdGNoZXMgcGVyZm9ybXMgcmVxdWVzdHMgaW4gYmF0Y2hlcwpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbkluQmF0Y2hlcyhzaXplIGludCwgc2xlZXAgaW50KSAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50IHsKCXAuY2xpZW50UlBDT3B0cyA9IGFwcGVuZChwLmNsaWVudFJQQ09wdHMsIHJwY2NsaWVudC5JbkJhdGNoZXMoc2l6ZSwgc2xlZXApKQoJcmV0dXJuIHAKfQoKLy8gT3B0aW9uRGlzY292ZXJ5VGltZW91dCBjb25maWd1cmVzIHRoZSByZXF1ZXN0IGRpc2NvdmVyeSB0aW1lb3V0LCBkZWZhdWx0cyB0byBjb25maWd1cmVkIGRpc2NvdmVyeSB0aW1lb3V0CmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uRGlzY292ZXJ5VGltZW91dCh0IHRpbWUuRHVyYXRpb24pICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5jbGllbnRSUENPcHRzID0gYXBwZW5kKHAuY2xpZW50UlBDT3B0cywgcnBjY2xpZW50LkRpc2NvdmVyeVRpbWVvdXQodCkpCglyZXR1cm4gcAp9CgovLyBPcHRpb25MaW1pdE1ldGhvZCBjb25maWd1cmVzIHRoZSBtZXRob2QgdG8gdXNlIHdoZW4gbGltaXRpbmcgdGFyZ2V0cyAtICJyYW5kb20iIG9yICJmaXJzdCIKZnVuYyAocCAqe3sgLkRETC5NZXRhZGF0YS5OYW1lIHwgU25ha2VUb0NhbWVsIH19Q2xpZW50KSBPcHRpb25MaW1pdE1ldGhvZChtIHN0cmluZykgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7CglwLmNsaWVudFJQQ09wdHMgPSBhcHBlbmQocC5jbGllbnRSUENPcHRzLCBycGNjbGllbnQuTGltaXRNZXRob2QobSkpCglyZXR1cm4gcAp9CgovLyBPcHRpb25MaW1pdFNpemUgc2V0cyBsaW1pdHMgb24gdGhlIHRhcmdldHMsIGVpdGhlciBhIG51bWJlciBvZiBhIHBlcmNlbnRhZ2UgbGlrZSAiMTAlIgpmdW5jIChwICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQpIE9wdGlvbkxpbWl0U2l6ZShzIHN0cmluZykgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCB7CglwLmNsaWVudFJQQ09wdHMgPSBhcHBlbmQocC5jbGllbnRSUENPcHRzLCBycGNjbGllbnQuTGltaXRTaXplKHMpKQoJcmV0dXJuIHAKfQoKLy8gT3B0aW9uTGltaXRTZWVkIHNldHMgdGhlIHJhbmRvbSBzZWVkIHVzZWQgdG8gc2VsZWN0IHRhcmdldHMgd2hlbiBsaW1pdGluZyBhbmQgbGltaXQgbWV0aG9kIGlzICJyYW5kb20iCmZ1bmMgKHAgKnt7IC5EREwuTWV0YWRhdGEuTmFtZSB8IFNuYWtlVG9DYW1lbCB9fUNsaWVudCkgT3B0aW9uTGltaXRTZWVkKHMgaW50NjQpICp7eyAuRERMLk1ldGFkYXRhLk5hbWUgfCBTbmFrZVRvQ2FtZWwgfX1DbGllbnQgewoJcC5jbGllbnRSUENPcHRzID0gYXBwZW5kKHAuY2xpZW50UlBDT3B0cywgcnBjY2xpZW50LkxpbWl0U2VlZChzKSkKCXJldHVybiBwCn0K",
}
