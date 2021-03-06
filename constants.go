package main

const axdTemplate string = `<?xml version="1.0" encoding="utf-8"?>
<appengine-web-app xmlns="http://appengine.google.com/ns/1.0">

    <!--must match project name-->
    <application>{{.Project}}</application>

    <!--always keep one until we implement rolling restarts etc-->
    <version>{{.SourceVersion}}</version>

    <!--change me, this forms the URL-->
    <module>{{.Module}}</module>

    <system-properties>
        <property name="com.acacia.source.outputpubsub" value="{{.TopicName}}"/>
        <property name="com.acacia.source.returnmessageids" value="true"/>
    </system-properties>
</appengine-web-app>
`
