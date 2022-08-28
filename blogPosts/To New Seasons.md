---
title: To New Seasons
date: August 21, 2022
tags: new season, next gen, golang, blog
series:
---

# To New Seasons

The site is finally live! This is my personal blog/thoughts collector. I wont ever write anything really personal here, as this is still the internet, but this will be more to talk about my thoughts on all things Tech around the world. The website will be in a beta state for a long time to come; I have a lot of ideas for the site, and it will take time to bring them together. It's like an evolving code project to show off my knowledge and learn.

The core of the site was already in place before I started designing too. I was really getting into backend server markdown parsers, and I started developing different backend rendered web sites around this concept. I have written many different builds in JS, TS, Nim, Goland and Rust. And I have bounced back and forth between tons of builds and frameworks and ideas... I have an issue with settling 😛. But for the most part Golang seems to fit all my needs and wants the best.

Rough idea of what I am looking for:

- Fast HTTP response times
- Good community support
- Well developed packages that support the project
- Semi-quick development
- Fun development 🎉

Golang just seems to fix this mold the best; I have written a lot in it as well so I feel comfortable and it's quick to pick up new ideas in and debug. Rust is a beautiful language with some amazing concepts, but in the HTTP space the performance is about the same vs Golang; which then makes it hard to justify the massive development time increase. Rust just makes you feel smart, it does things 'right' and forces you to aswell... it just takes a long time to develop in. And when learning new concepts in Rust you are left trying to wrap your mind around these huge boilerplate concepts and large complex types. I just want to write code, see it come to life and have fun! If you're looking for a good Rust HTTP server, start with Axum! It is probably the best Rust server you can use right now! Great project by the Tokio team.

And then there is the JS/TS world, which I hate the state of it right now. Give it a couple more years and it will be in a perfect state. Nodejs should probably just die at this point, and I am only slightly joking. Node is just missing a lot when it comes to a user friendly modern JS/TS experience; go try and use top level async/await with TypeScript on a Node project... good luck! And then Deno offers great modern tooling, but performance is still behind and packages are minimal. Deno has a great potential but is just too young; they are having a hard time convincing people to move over from Node and loose 95% of the mature packages they use everyday. And then you have Bun, which is super fast but again very very early and just not ready yet. Give it a couple years and I think Deno and Bun take off. Node will probably be around forever just to maintain legacy code; but a lot of developers have jumped ship to Golang/Rust already.

And then the question is "what framework should I use.. or should I even use a framework?". Yeah use a framework... some people are purists in the Golang space and will only use http/net. That is awesome and fine; you will learn a ton about how HTTP works behind all the magic, but you will spend a long time probably doing things wrong. Pick a good framework or micro-framework and start there. I do think Go's http/net package is great and made and maintainted by really smart people, so I picked a framework built around it: Gin!

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    app := gin.Default()

    http.ListenAndServe(HTTP_PORT), app)
}
```

Here is a little fun code snippet of a Gin server! Good stuff!! If you are new to coding or coming from Nodejs, take a look at <a href="https://github.com/gofiber/fiber" target="_blank" rel="noreferrer noopener">GoFiber</a>. Many people in the Golang world will encourage you not to use GoFiber (as it is not built on http/net), but it is a good place to start learning HTTP Golang. Once you feel comfortable there, switch to a http/net based framework.🚀

So what is next?! Well hopefully I keep developing this site and learning new things to implement. And I hope the content is good and enjoyable!

### Eric Christensen
