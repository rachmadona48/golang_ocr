import PyPDF2
# The opening procedure for a file object variable will be rb
pdffile = open(r'sample.pdf', 'rb')
# create a variable called reader that will read the pdf file
pdfReader = PyPDF2.PdfFileReader(pdffile)
# The number of pages in this pdf file will be saved.
num = pdfReader.numPages
print(num)
# create a variable that will select the selected number of pages
pageobj = pdfReader.getPage(0)
resulttext = pageobj.extractText()
newfile = open(
    r"sample.txt", "a")
newfile.writelines(resulttext)