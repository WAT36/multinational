const log4js = require('log4js');

log4js.configure({
    appenders: {
        out:{
            type: "dateFile",
            filename: "all-the-logs.log",
            pattern: "-yyyy-MM-dd",
            numBackups: 5
        }
    },
    categories: {
        default: {
            appenders: ["out"],
            level: "info",
        }
    }
});

const logger = log4js.getLogger();
logger.info("I will be logged this as info in all-the-logs.log");