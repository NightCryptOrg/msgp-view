# Installation vars
PREFIX ?= /usr/local
DATADIR ?= $(PREFIX)/share

# Build using go build

install: msgp-view
	install -d $(DESTDIR)$(PREFIX)/bin
	install msgp-view $(DESTDIR)$(PREFIX)/bin/msgp-view
	install -d $(DESTDIR)$(DATADIR)/doc/msgp-view
	install -m644 README.md $(DESTDIR)$(DATADIR)/doc/msgp-view/README.md
	install -d $(DESTDIR)$(DATADIR)/licenses/msgp-view
	install -m644 LICENSE $(DESTDIR)$(DATADIR)/licenses/msgp-view/LICENSE
.PHONY: install

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/msgp-view
	rm -rf $(DESTDIR)$(DATADIR)/doc/msgp-view
	rm -rf $(DESTDIR)$(DATADIR)/licenses/msgp-view
.PHONY: uninstall
