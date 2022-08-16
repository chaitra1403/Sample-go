import logging
import logging.handlers as handlers
import time

lgr = logging.getLogger('Roostai')
lgr.setLevel(logging.INFO)

logHlr = handlers.TimedRotatingFileHandler('app.log', when='M', interval=1)
logHlr.setLevel(logging.INFO)
lgr.addHandler(logHlr)

def main():
    while True:
        time.sleep(1)
        lgr.info("Hello World")

main()
