import logging
import logging.handlers as handlers
import time

lgr = logging.getLogger('Roostai')
lgr.setLevel(logging.INFO)

logHlr = handlers.RotatingFileHandler('app.log', maxBytes=500, backupCount=2)
logHlr.setLevel(logging.INFO)
lgr.addHandler(logHlr)

def main():
    while True:
        time.sleep(1)
        lgr.info("Hello World")

main()
