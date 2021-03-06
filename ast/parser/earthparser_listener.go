// Code generated from ast/parser/EarthParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // EarthParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EarthParserListener is a complete listener for a parse tree produced by EarthParser.
type EarthParserListener interface {
	antlr.ParseTreeListener

	// EnterEarthFile is called when entering the earthFile production.
	EnterEarthFile(c *EarthFileContext)

	// EnterTargets is called when entering the targets production.
	EnterTargets(c *TargetsContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterTargetHeader is called when entering the targetHeader production.
	EnterTargetHeader(c *TargetHeaderContext)

	// EnterStmts is called when entering the stmts production.
	EnterStmts(c *StmtsContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterCommandStmt is called when entering the commandStmt production.
	EnterCommandStmt(c *CommandStmtContext)

	// EnterWithStmt is called when entering the withStmt production.
	EnterWithStmt(c *WithStmtContext)

	// EnterWithBlock is called when entering the withBlock production.
	EnterWithBlock(c *WithBlockContext)

	// EnterWithExpr is called when entering the withExpr production.
	EnterWithExpr(c *WithExprContext)

	// EnterWithCommand is called when entering the withCommand production.
	EnterWithCommand(c *WithCommandContext)

	// EnterDockerCommand is called when entering the dockerCommand production.
	EnterDockerCommand(c *DockerCommandContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterIfClause is called when entering the ifClause production.
	EnterIfClause(c *IfClauseContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterElseIfClause is called when entering the elseIfClause production.
	EnterElseIfClause(c *ElseIfClauseContext)

	// EnterElseIfBlock is called when entering the elseIfBlock production.
	EnterElseIfBlock(c *ElseIfBlockContext)

	// EnterElseClause is called when entering the elseClause production.
	EnterElseClause(c *ElseClauseContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterElseIfExpr is called when entering the elseIfExpr production.
	EnterElseIfExpr(c *ElseIfExprContext)

	// EnterFromStmt is called when entering the fromStmt production.
	EnterFromStmt(c *FromStmtContext)

	// EnterFromDockerfileStmt is called when entering the fromDockerfileStmt production.
	EnterFromDockerfileStmt(c *FromDockerfileStmtContext)

	// EnterLocallyStmt is called when entering the locallyStmt production.
	EnterLocallyStmt(c *LocallyStmtContext)

	// EnterCopyStmt is called when entering the copyStmt production.
	EnterCopyStmt(c *CopyStmtContext)

	// EnterSaveStmt is called when entering the saveStmt production.
	EnterSaveStmt(c *SaveStmtContext)

	// EnterSaveImage is called when entering the saveImage production.
	EnterSaveImage(c *SaveImageContext)

	// EnterSaveArtifact is called when entering the saveArtifact production.
	EnterSaveArtifact(c *SaveArtifactContext)

	// EnterRunStmt is called when entering the runStmt production.
	EnterRunStmt(c *RunStmtContext)

	// EnterBuildStmt is called when entering the buildStmt production.
	EnterBuildStmt(c *BuildStmtContext)

	// EnterWorkdirStmt is called when entering the workdirStmt production.
	EnterWorkdirStmt(c *WorkdirStmtContext)

	// EnterUserStmt is called when entering the userStmt production.
	EnterUserStmt(c *UserStmtContext)

	// EnterCmdStmt is called when entering the cmdStmt production.
	EnterCmdStmt(c *CmdStmtContext)

	// EnterEntrypointStmt is called when entering the entrypointStmt production.
	EnterEntrypointStmt(c *EntrypointStmtContext)

	// EnterExposeStmt is called when entering the exposeStmt production.
	EnterExposeStmt(c *ExposeStmtContext)

	// EnterVolumeStmt is called when entering the volumeStmt production.
	EnterVolumeStmt(c *VolumeStmtContext)

	// EnterEnvStmt is called when entering the envStmt production.
	EnterEnvStmt(c *EnvStmtContext)

	// EnterArgStmt is called when entering the argStmt production.
	EnterArgStmt(c *ArgStmtContext)

	// EnterEnvArgKey is called when entering the envArgKey production.
	EnterEnvArgKey(c *EnvArgKeyContext)

	// EnterEnvArgValue is called when entering the envArgValue production.
	EnterEnvArgValue(c *EnvArgValueContext)

	// EnterLabelStmt is called when entering the labelStmt production.
	EnterLabelStmt(c *LabelStmtContext)

	// EnterLabelKey is called when entering the labelKey production.
	EnterLabelKey(c *LabelKeyContext)

	// EnterLabelValue is called when entering the labelValue production.
	EnterLabelValue(c *LabelValueContext)

	// EnterGitCloneStmt is called when entering the gitCloneStmt production.
	EnterGitCloneStmt(c *GitCloneStmtContext)

	// EnterAddStmt is called when entering the addStmt production.
	EnterAddStmt(c *AddStmtContext)

	// EnterStopsignalStmt is called when entering the stopsignalStmt production.
	EnterStopsignalStmt(c *StopsignalStmtContext)

	// EnterOnbuildStmt is called when entering the onbuildStmt production.
	EnterOnbuildStmt(c *OnbuildStmtContext)

	// EnterHealthcheckStmt is called when entering the healthcheckStmt production.
	EnterHealthcheckStmt(c *HealthcheckStmtContext)

	// EnterShellStmt is called when entering the shellStmt production.
	EnterShellStmt(c *ShellStmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterStmtWordsMaybeJSON is called when entering the stmtWordsMaybeJSON production.
	EnterStmtWordsMaybeJSON(c *StmtWordsMaybeJSONContext)

	// EnterStmtWords is called when entering the stmtWords production.
	EnterStmtWords(c *StmtWordsContext)

	// EnterStmtWord is called when entering the stmtWord production.
	EnterStmtWord(c *StmtWordContext)

	// ExitEarthFile is called when exiting the earthFile production.
	ExitEarthFile(c *EarthFileContext)

	// ExitTargets is called when exiting the targets production.
	ExitTargets(c *TargetsContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitTargetHeader is called when exiting the targetHeader production.
	ExitTargetHeader(c *TargetHeaderContext)

	// ExitStmts is called when exiting the stmts production.
	ExitStmts(c *StmtsContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitCommandStmt is called when exiting the commandStmt production.
	ExitCommandStmt(c *CommandStmtContext)

	// ExitWithStmt is called when exiting the withStmt production.
	ExitWithStmt(c *WithStmtContext)

	// ExitWithBlock is called when exiting the withBlock production.
	ExitWithBlock(c *WithBlockContext)

	// ExitWithExpr is called when exiting the withExpr production.
	ExitWithExpr(c *WithExprContext)

	// ExitWithCommand is called when exiting the withCommand production.
	ExitWithCommand(c *WithCommandContext)

	// ExitDockerCommand is called when exiting the dockerCommand production.
	ExitDockerCommand(c *DockerCommandContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitIfClause is called when exiting the ifClause production.
	ExitIfClause(c *IfClauseContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitElseIfClause is called when exiting the elseIfClause production.
	ExitElseIfClause(c *ElseIfClauseContext)

	// ExitElseIfBlock is called when exiting the elseIfBlock production.
	ExitElseIfBlock(c *ElseIfBlockContext)

	// ExitElseClause is called when exiting the elseClause production.
	ExitElseClause(c *ElseClauseContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitElseIfExpr is called when exiting the elseIfExpr production.
	ExitElseIfExpr(c *ElseIfExprContext)

	// ExitFromStmt is called when exiting the fromStmt production.
	ExitFromStmt(c *FromStmtContext)

	// ExitFromDockerfileStmt is called when exiting the fromDockerfileStmt production.
	ExitFromDockerfileStmt(c *FromDockerfileStmtContext)

	// ExitLocallyStmt is called when exiting the locallyStmt production.
	ExitLocallyStmt(c *LocallyStmtContext)

	// ExitCopyStmt is called when exiting the copyStmt production.
	ExitCopyStmt(c *CopyStmtContext)

	// ExitSaveStmt is called when exiting the saveStmt production.
	ExitSaveStmt(c *SaveStmtContext)

	// ExitSaveImage is called when exiting the saveImage production.
	ExitSaveImage(c *SaveImageContext)

	// ExitSaveArtifact is called when exiting the saveArtifact production.
	ExitSaveArtifact(c *SaveArtifactContext)

	// ExitRunStmt is called when exiting the runStmt production.
	ExitRunStmt(c *RunStmtContext)

	// ExitBuildStmt is called when exiting the buildStmt production.
	ExitBuildStmt(c *BuildStmtContext)

	// ExitWorkdirStmt is called when exiting the workdirStmt production.
	ExitWorkdirStmt(c *WorkdirStmtContext)

	// ExitUserStmt is called when exiting the userStmt production.
	ExitUserStmt(c *UserStmtContext)

	// ExitCmdStmt is called when exiting the cmdStmt production.
	ExitCmdStmt(c *CmdStmtContext)

	// ExitEntrypointStmt is called when exiting the entrypointStmt production.
	ExitEntrypointStmt(c *EntrypointStmtContext)

	// ExitExposeStmt is called when exiting the exposeStmt production.
	ExitExposeStmt(c *ExposeStmtContext)

	// ExitVolumeStmt is called when exiting the volumeStmt production.
	ExitVolumeStmt(c *VolumeStmtContext)

	// ExitEnvStmt is called when exiting the envStmt production.
	ExitEnvStmt(c *EnvStmtContext)

	// ExitArgStmt is called when exiting the argStmt production.
	ExitArgStmt(c *ArgStmtContext)

	// ExitEnvArgKey is called when exiting the envArgKey production.
	ExitEnvArgKey(c *EnvArgKeyContext)

	// ExitEnvArgValue is called when exiting the envArgValue production.
	ExitEnvArgValue(c *EnvArgValueContext)

	// ExitLabelStmt is called when exiting the labelStmt production.
	ExitLabelStmt(c *LabelStmtContext)

	// ExitLabelKey is called when exiting the labelKey production.
	ExitLabelKey(c *LabelKeyContext)

	// ExitLabelValue is called when exiting the labelValue production.
	ExitLabelValue(c *LabelValueContext)

	// ExitGitCloneStmt is called when exiting the gitCloneStmt production.
	ExitGitCloneStmt(c *GitCloneStmtContext)

	// ExitAddStmt is called when exiting the addStmt production.
	ExitAddStmt(c *AddStmtContext)

	// ExitStopsignalStmt is called when exiting the stopsignalStmt production.
	ExitStopsignalStmt(c *StopsignalStmtContext)

	// ExitOnbuildStmt is called when exiting the onbuildStmt production.
	ExitOnbuildStmt(c *OnbuildStmtContext)

	// ExitHealthcheckStmt is called when exiting the healthcheckStmt production.
	ExitHealthcheckStmt(c *HealthcheckStmtContext)

	// ExitShellStmt is called when exiting the shellStmt production.
	ExitShellStmt(c *ShellStmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitStmtWordsMaybeJSON is called when exiting the stmtWordsMaybeJSON production.
	ExitStmtWordsMaybeJSON(c *StmtWordsMaybeJSONContext)

	// ExitStmtWords is called when exiting the stmtWords production.
	ExitStmtWords(c *StmtWordsContext)

	// ExitStmtWord is called when exiting the stmtWord production.
	ExitStmtWord(c *StmtWordContext)
}
