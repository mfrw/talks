all: hsa.pdf

hsa.pdf: $(wildcard *.tex)
	pdflatex hsa.tex

clean:
	-rm -f *.aux
	-rm -f *.bbl
	-rm -f *.blg
	-rm -f *.log
	-rm -f *.out
	-rm -f hsa.pdf
