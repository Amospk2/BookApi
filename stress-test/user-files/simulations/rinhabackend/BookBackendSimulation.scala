import scala.concurrent.duration._

import scala.util.Random

import io.gatling.core.Predef._
import io.gatling.http.Predef._


class RinhaBackendSimulation
  extends Simulation {

  val httpProtocol = http
    .baseUrl("http://localhost:9999")
    .userAgentHeader("Agente do Caos - 2023")

  val criacaoEConsultaAutores = scenario("Criação E Talvez Consulta de Autores")
    .feed(tsv("authors-payloads.tsv").circular())
    .exec(
      http("criação")
      .post("/pessoas").body(StringBody("#{payload}"))
      .header("content-type", "application/json")
      .check(status.in(201, 422, 400))
      .check(status.saveAs("httpStatus"))
      .checkIf(session => session("httpStatus").as[String] == "201") {
        header("Location").saveAs("location")
      }
    )
    .pause(1.milliseconds, 30.milliseconds)
    .doIf(session => session.contains("location")) {
      exec(
        http("consulta")
        .get("#{location}")
      )
    }


  val criacaoInvalida = scenario("Busca Inválida de Pessoas")
    .exec(
      http("busca inválida")
      .post("/pessoas").body(StringBody("#{payload}"))
      .header("content-type", "application/json")
      // 400 - bad request se não passar 't' como query string
      .check(status.is(400))
    )


  setUp(
    criacaoEConsultaAutores.inject(
      constantUsersPerSec(2).during(10.seconds), // warm up
      constantUsersPerSec(5).during(15.seconds).randomized, // are you ready?
      
      rampUsersPerSec(6).to(600).during(3.minutes) // lezzz go!!!
    ),
    criacaoInvalida.inject(
      constantUsersPerSec(2).during(25.seconds), // warm up
      
      rampUsersPerSec(6).to(40).during(3.minutes) // lezzz go!!!
    )

  ).protocols(httpProtocol)
}
