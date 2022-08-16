import logging
logging.basicConfig(filename='roostlog.log', encoding='utf-8', level=logging.DEBUG)
logging.debug('This is debug log')
logging.info('This is Info log')
logging.warning('This is warning log')
logging.error('And this is error log')
