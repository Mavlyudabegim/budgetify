require('dotenv').config();
const cors = require('cors');
const express = require('express');
const accountRouter = require('./routes/account');
const userRouter = require('./routes/user');
const categoryRouter = require('./routes/category');
const transactionRouter = require('./routes/transaction');
const subscriptionRouter = require('./routes/subscription');
const piggybankRouter = require('./routes/piggybank');
const errorMiddleware = require('./middleware/error-middleware');
const app = express();
const cookieParser = require('cookie-parser');
app.use(cors());
app.use(cookieParser());
app.use('/api', userRouter);
app.use('/api/accounts', accountRouter);
app.use('/api/transactions', transactionRouter);
app.use('/api/categories', categoryRouter);
app.use('/api/subscriptions', subscriptionRouter);
app.use('/api/piggybanks', piggybankRouter);
app.use(express.urlencoded({ extended: false }));
app.use(express.static('public'));
app.use(errorMiddleware);
module.exports = app;
