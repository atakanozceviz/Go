<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0"
xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
  <xsl:template match="/cmf">
    <html lang="tr">
    <head>
    <title>Öğrenci Bilgileri</title>
		<link href="css/bootstrap.min.css" rel="stylesheet"></link>
		<link href="css/main.css" rel="stylesheet"></link>
    </head>
    
    <body>
    <nav class="navbar navbar-inverse navbar-static-top">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar"> <span class="sr-only">Toggle navigation</span> <span class="icon-bar"></span> <span class="icon-bar"></span> <span class="icon-bar"></span> </button>
          <a class="navbar-brand" href="#">Ödev-1</a> </div>
        <div id="navbar" class="collapse navbar-collapse">
          <ul class="nav navbar-nav">
            <li><a href="index.html">Kişisel Bilgiler</a></li>
            <li class="active"><a href="cmf.xml">Öğrenci Bilgileri</a></li>
            <li><a href="json.html">JSON</a></li>
            <li><a href="form.html">Form</a></li>
          </ul>
        </div>
        <!--/.nav-collapse --> 
      </div>
    </nav>
    <div class="container">
      <div class="row">
        <div class="col-xs-12">
          <h2>Öğrenciler</h2>
          <table class="table table-bordered">
            <thead>
              <th>Ad</th>
              <th>Soyad</th>
              <th>Devamsızlık</th>
              <th>Not</th>
            </thead>
			  <tbody>
            <xsl:for-each select="ogrenci">
              <tr>
                <td><xsl:value-of select="ad"/></td>
                <td><xsl:value-of select="soyad"/></td>
                <xsl:choose>
                  <xsl:when test="devamsizlik &gt; 3">
                    <td bgcolor="#5882FA">Devamsızlıktan Kaldı(
                      <xsl:value-of select="devamsizlik"/>
                      )</td>
                  </xsl:when>
                  <xsl:otherwise>
                    <td><xsl:value-of select="devamsizlik"/></td>
                  </xsl:otherwise>
                </xsl:choose>
                <xsl:choose>
                  <xsl:when test="not &lt; 60">
                    <td bgcolor="#A9F5BC">Sınavdan Kaldı(
                      <xsl:value-of select="not"/>
                      )</td>
                  </xsl:when>
                  <xsl:otherwise>
                    <td><xsl:value-of select="not"/></td>
                  </xsl:otherwise>
                </xsl:choose>
              </tr>
            </xsl:for-each>
            </tbody>
          </table>
        </div>
      </div>
    </div>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script> 
    <script src="js/bootstrap.min.js"></script>
    </body>
    </html>
  </xsl:template>
</xsl:stylesheet>