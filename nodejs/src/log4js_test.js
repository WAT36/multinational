const log4js = require('log4js');

// 通常時
const logger = log4js.getLogger()
logger.level = 'all'
logger.info('info test messages')

// 別のロガーオブジェクト定義
const cheese = log4js.getLogger('cheese')
cheese.level = 'all'
cheese.info('info cheese fondu')