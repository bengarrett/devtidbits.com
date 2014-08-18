<!----- xml-high-safe.cfc 1.0 by Ben Garrett : 11 March 2008 :
Released under the Creative Commons BSD License (http://creativecommons.org/licenses/BSD/) ----->
<cfcomponent output="no">
	<cffunction name="XMLHighSafe" access="public" returntype="string" output="no" hint="This scans through a string, finds any characters that have a higher ASCII numeric value greater than 127 and automatically convert them to a hexadecimal numeric character">
		<cfargument name="text" type="string" required="yes">
		<cfscript>
			var i = 0;
			var tmp = '';
			while(ReFind('[^\x00-\x7F]',text,i,false))
			{
				i = ReFind('[^\x00-\x7F]',text,i,false); // discover high chr and save it's numeric string position.
				tmp = '&##x#FormatBaseN(Asc(Mid(text,i,1)),16)#;'; // obtain the high chr and convert it to a hex numeric chr.
				text = Insert(tmp,text,i); // insert the new hex numeric chr into the string.
				text = RemoveChars(text,i,1); // delete the redundant high chr from string.
				i = i+Len(tmp); // adjust the loop scan for the new chr placement, then continue the loop.
			}
			return text;
		</cfscript>
	</cffunction>
</cfcomponent>