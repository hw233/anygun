########################################################################
# bintalk compiler
########################################################################

# find bison & flex.
find_package(BISON REQUIRED)
find_package(FLEX REQUIRED)
BISON_TARGET(BINTALK_PARSER bin.y ${CMAKE_CURRENT_BINARY_DIR}/bin.tab.cpp)
IF(MSVC)
	FLEX_TARGET(BINTALK_SCANNER bin.l ${CMAKE_CURRENT_BINARY_DIR}/lex.yy.cpp COMPILE_FLAGS "--nounistd" )
ELSE()
	FLEX_TARGET(BINTALK_SCANNER bin.l ${CMAKE_CURRENT_BINARY_DIR}/lex.yy.cpp)
ENDIF()
ADD_FLEX_BISON_DEPENDENCY(BINTALK_SCANNER BINTALK_PARSER)

#SET(Boost_USE_STATIC_LIBS ON)
#FIND_PACKAGE(Boost COMPONENTS system filesystem program_options REQUIRED)
#INCLUDE_DIRECTORIES( ${Boost_INCLUDE_DIR} )
#LINK_DIRECTORIES(${Boost_LIBRARY_DIR})


# to find generated files.
include_directories(${CMAKE_CURRENT_SOURCE_DIR})
# make generated files include source files.
include_directories(${CMAKE_CURRENT_BINARY_DIR})

# sources.
set(PARSER_FILES
	${BISON_BINTALK_PARSER_OUTPUTS}
	${FLEX_BINTALK_SCANNER_OUTPUTS}
	bin.y
	bin.l
	)
set(SPEC_FILES
	Definition.h
	Enum.h
	Field.h
	Service.h
	Struct.h
	Enum.cpp
	Field.cpp
	Service.cpp
	Struct.cpp
	
	)
set(GEN_FILES
	CppGenerator.h
	PYGenerator.h
	ERLGenerator.h
	CSGenerator.h
#    ASGenerator.h

	CppGenerator.cpp
	PYGenerator.cpp
	ERLGenerator.cpp
	CSGenerator.cpp
 #   ASGenerator.cpp
	)
set(MISC_FILES
	CodeFile.h
	CodeFile.cpp
	CodeGenerator.h
    Compiler.h
    Compiler.cpp
	Args.h
	Args.cpp
    Main.cpp
	)

source_group( parser FILES ${PARSER_FILES})
source_group( spec FILES ${SPEC_FILES})
source_group( gen FILES ${GEN_FILES})
source_group( misc FILES ${MISC_FILES})

add_executable(rpc ${PARSER_FILES} ${SPEC_FILES} ${GEN_FILES} ${MISC_FILES})

# install bintalk
#TARGET_LINK_LIBRARIES(rpc ${Boost_LIBRARIES})
install(TARGETS rpc DESTINATION bin)
