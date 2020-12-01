FROM maven:3.6.3-openjdk-14-slim AS build

COPY settings.xml /usr/share/maven/conf/

COPY pom.xml pom.xml
COPY pos-api/pom.xml pos-api/pom.xml
COPY pos-model/pom.xml pos-model/pom.xml
COPY pos-base/pom.xml pos-base/pom.xml
COPY pos-script/pom.xml pos-script/pom.xml

RUN mvn dependency:go-offline package -B

COPY pos-api/src pos-api/src
COPY pos-model/src pos-model/src
COPY pos-base/src pos-base/src
COPY pos-script/src pos-script/src

RUN mvn install

FROM groovy:3.0.5-jdk14
USER root

WORKDIR /

RUN mkdir service

COPY --from=build /pos-base/target/ /service/
COPY /pos-script/src/main/groovy/com/atlas/pos/script/portal /service/script/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

ENV JAVA_TOOL_OPTIONS -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:5005

EXPOSE 5005

CMD /wait && java --enable-preview -jar /service/pos-base-1.0-SNAPSHOT.jar -Xdebug