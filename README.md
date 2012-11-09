Ops School Curriculum
=====================
[![Build Status](https://secure.travis-ci.org/opsschool/curriculum.png?branch=master)](https://travis-ci.org/opsschool/curriculum)

The current documentation based on these sources can be seen at:
http://www.opsschool.org/

Welcome!

If you have arrived here, you are probably interested in helping out. So thank you for your time.

Things you should know:

* This project is written in [reStructuredText](http://docutils.sourceforge.net/docs/user/rst/quickstart.html)
* Hosted by [Read the Docs](http://readthedocs.org/)
* Tested by rendering in [Sphinx](http://sphinx-doc.org/) on [Travis CI](https://travis-ci.org)

This is the only Markdown file in the repository, as it's not meant to be included in the documentation itself.

If you are looking to add content, fix formatting, syntax, typos or other wonderful things, please follow this process:

* Install Sphinx: `easy_install Sphinx` or `pip install Sphinx`
* Fork the `opsschool/curriculum` repo to your own account
* Check out a branch to make your changes on: `git checkout --branch <my_topic>`
* Execute `make html` to build the docs in to `_build/`
* Make your changes
* Execute `make html` again and verify your changes don't cause any warnings/errors
* Commit with a descriptive message, and submit a pull request from your branch to `master`
* One of the editors will review the change, and either merge it or provide some feedback. Community review is also encouraged.

Some cool things:

* `vim-common` contains a reStructuredText syntax highlighter
* The [Emacs support](http://docutils.sourceforge.net/docs/user/emacs.html) via rst-mode comes as part of the docutils package under /docutils/tools/editors/emacs/rst.el
