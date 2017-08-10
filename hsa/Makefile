all: hsa.pdf

hsa.pdf: $(wildcard *.tex) hsa.bib
	pdflatex hsa.tex
	bibtex hsa
	pdflatex hsa.tex
	pdflatex hsa.tex

clean:
	-rm -f *.aux
	-rm -f *.bbl
	-rm -f *.blg
	-rm -f *.log
	-rm -f *.out
	-rm -f hsa.pdf
