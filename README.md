# AgentReport

<br>

## Introduction

Hey there are you also tired of keeping up with news articles and cross checking different news sources to see if you are getting the correct news or not. Well now with this new AI agent you won't have to worry about anything.

Agent Report scrapes the internet and the major news outlets such as [CNN](https://edition.cnn.com/), [BBC](https://www.bbc.com/) and many more to get news articles. An AI voice agent that uses [Vapi](https://vapi.ai) as a virtual assistant will then be ready to answer all your questions about the current affairs and even summarize the today's news or any previous news which you would like.

<br>

## System Design

Agent Report follows a workflow that is similar to the one given below:

<img src="./public/sysdesign.jpg" alt="System Desgin">

<br>

## Backend Structure

The backend is responsible for Scraping the data from the sites. It exposes an API endpoint which is used by the frontend to send requests by the user and give feedback via Vapi agent. The backend is written entirely in Go.

<img src="./public/struct.jpg" alt="Backend Structure">
